package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// simple in-memory user DB
type User struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

var users = map[string]User{}

var oauthConfig *oauth2.Config
var jwtKey []byte

func init() {
	oauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"openid", "profile", "email"},
		RedirectURL:  os.Getenv("BACKEND_URL") + "/auth/google/callback",
	}
	jwtKey = []byte(os.Getenv("SESSION_KEY"))
}

// GET /auth/google/login
func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	state := fmt.Sprintf("%d", time.Now().UnixNano())
	// optional: store state in cookie/session to validate in callback
	url := oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

// GET /auth/google/callback
func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "code not found", http.StatusBadRequest)
		return
	}

	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "token exchange failed", http.StatusInternalServerError)
		log.Println("exchange:", err)
		return
	}

	// fetch userinfo
	client := oauthConfig.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "failed to get userinfo", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var gi struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}
	if err := json.Unmarshal(body, &gi); err != nil {
		http.Error(w, "failed to parse userinfo", http.StatusInternalServerError)
		return
	}

	// create-or-update user in memory (replace with DB code)
	user := User{ID: gi.ID, Email: gi.Email, Name: gi.Name, Picture: gi.Picture}
	users[user.ID] = user

	// create JWT session
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := jwtToken.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "failed to sign token", http.StatusInternalServerError)
		return
	}

	// set cookie (HttpOnly, Secure in prod)
	cookie := &http.Cookie{
		Name:     "session",
		Value:    signed,
		Path:     "/",
		HttpOnly: true,
		Secure:   os.Getenv("BACKEND_URL") != "http://localhost:8080",
		SameSite: http.SameSiteNoneMode,
		MaxAge:   7 * 24 * 3600,
	}
	http.SetCookie(w, cookie)

	// redirect back to frontend
	frontend := os.Getenv("FRONTEND_URL")
	if frontend == "" {
		frontend = "/"
	}
	http.Redirect(w, r, frontend, http.StatusFound)
}

// GET /api/me
func Me(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tokenStr := c.Value

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	uid := claims["sub"].(string)
	user, ok := users[uid]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(user)
}
