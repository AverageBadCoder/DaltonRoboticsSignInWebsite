package main

import (
	"log"
	"net/http"
	"os"

	"my-go-backend/internal/routes"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	// load .env (optional in prod where env vars are provided by the host)
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := routes.SetupRoutes()

	frontend := os.Getenv("FRONTEND_URL")
	ngrok := "https://unmordantly-stirruplike-naida.ngrok-free.dev"
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{frontend, "http://localhost:5173", ngrok}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, cors(router)))
}
