package handlers

import (
	"encoding/json"
	"net/http"
)

// Response represents the structure of the response message
type Response struct {
	Message string `json:"message"`
}

// HelloWorld is a simple handler that responds with a hello world message
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

// ApiHandler is a simple example endpoint used by routes.SetupRoutes.
func ApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]interface{}{
		"ok":      true,
		"message": "api endpoint working",
		"method":  r.Method,
	}
	_ = json.NewEncoder(w).Encode(resp)
}
