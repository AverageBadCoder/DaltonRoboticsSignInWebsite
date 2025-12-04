package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// simple in-memory user DB
type User struct {
	ID       string `json:"id"`       //username
	Email    string `json:"email"`    //email
	Name     string `json:"name"`     //full name of user
	Picture  string `json:"picture"`  //profile picture
	Tag      string `json:"tag"`      //admin or team tags will descide the access level this must be set by an existing admin user
	Approval bool   `json:"approval"` //approval status must be set by an existing admin user before user or admin can acsess website
}

var users = map[string]User{}

var oauthConfig *oauth2.Config
var jwtKey []byte

func init() {
	// ensure .env is loaded before reading env vars (main's godotenv.Load() is too late
	// because package init runs earlier)
	_ = godotenv.Load()

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	// fallback: try credentials.json if needed (existing logic can remain)
	if clientID == "" || clientSecret == "" {
		log.Println("GOOGLE_CLIENT_ID or GOOGLE_CLIENT_SECRET not set, falling back to credentials.json")
		b, err := os.ReadFile("credentials.json")
		if err != nil {
			log.Fatalf("failed to read credentials.json: %v", err)
		}
		var cred struct {
			ClientID     string `json:"client_id"`
			ClientSecret string `json:"client_secret"`
		}
		if err := json.Unmarshal(b, &cred); err != nil {
			log.Fatalf("failed to parse credentials.json: %v", err)
		}
		clientID = cred.ClientID
		clientSecret = cred.ClientSecret
	}

	oauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{"openid", "profile", "email"},
		RedirectURL:  os.Getenv("BACKEND_URL") + "/auth/google/callback",
	}

	jwtKey = []byte(os.Getenv("SESSION_KEY"))

	log.Printf("oauth client id=%q redirect_url=%q\n", oauthConfig.ClientID, oauthConfig.RedirectURL)
}

// GET /auth/google/login
func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	state := fmt.Sprintf("%d", time.Now().UnixNano())
	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
	})
	url := oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	log.Printf("GoogleLogin redirecting to: %s (state=%s)\n", url, state)
	http.Redirect(w, r, url, http.StatusFound)
}

// GET /auth/google/callback
func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	log.Printf("GoogleCallback raw query: %s\n", r.URL.RawQuery)
	if errParam := r.URL.Query().Get("error"); errParam != "" {
		log.Printf("Google callback returned error=%s\n", errParam)
		http.Error(w, "oauth error: "+errParam, http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "code not found", http.StatusBadRequest)
		return
	}

	// try normal exchange first
	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		// detailed logging
		log.Printf("oauth exchange error: %v (client_id=%q redirect=%q code=%q)\n", err, oauthConfig.ClientID, oauthConfig.RedirectURL, code)

		// manual token request to capture provider response body for debugging
		resp, respErr := http.PostForm("https://oauth2.googleapis.com/token", url.Values{
			"code":          {code},
			"client_id":     {oauthConfig.ClientID},
			"client_secret": {oauthConfig.ClientSecret},
			"redirect_uri":  {oauthConfig.RedirectURL},
			"grant_type":    {"authorization_code"},
		})
		if respErr != nil {
			log.Printf("manual token request failed: %v\n", respErr)
			http.Error(w, "token exchange failed", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		b, _ := io.ReadAll(resp.Body)
		log.Printf("manual token response status=%d body=%s\n", resp.StatusCode, string(b))

		http.Error(w, "token exchange failed", http.StatusInternalServerError)
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
		ID       string `json:"id"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Picture  string `json:"picture"`
		Tag      string `json:"tag"`
		Approval bool   `json:"approval"`
	}
	if err := json.Unmarshal(body, &gi); err != nil {
		http.Error(w, "failed to parse userinfo", http.StatusInternalServerError)
		return
	}

	// create-or-update user in memory (replace with DB code)
	user := User{ID: gi.ID, Email: gi.Email, Name: gi.Name, Picture: gi.Picture, Tag: gi.Tag, Approval: gi.Approval}
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

	// redirect to frontend with token in fragment so SPA can read it
	frontend := os.Getenv("FRONTEND_URL")
	if frontend == "" {
		frontend = "/"
	}
	// ensure no trailing slash
	frontend = strings.TrimRight(frontend, "/")
	http.Redirect(w, r, frontend+"/#token="+signed, http.StatusFound)

}

// GET /api/me
func Me(w http.ResponseWriter, r *http.Request) {
	// 1) Try Authorization header Bearer <token>
	auth := r.Header.Get("Authorization")
	var tokenStr string
	if strings.HasPrefix(auth, "Bearer ") {
		tokenStr = strings.TrimPrefix(auth, "Bearer ")
	} else {
		// fallback to cookie if present
		c, err := r.Cookie("session")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenStr = c.Value
	}

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
	uid, _ := claims["sub"].(string)
	user, ok := users[uid]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
