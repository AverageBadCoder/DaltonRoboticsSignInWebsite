package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"my-go-backend/internal/routes"
)

func main() {
	r := mux.NewRouter()
	routes.SetupRoutes(r)

	http.Handle("/", r)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}