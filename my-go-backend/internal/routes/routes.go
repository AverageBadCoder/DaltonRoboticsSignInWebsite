package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"my-go-backend/internal/handlers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/endpoint", handlers.ApiHandler).Methods(http.MethodGet, http.MethodPost)

	return router
}