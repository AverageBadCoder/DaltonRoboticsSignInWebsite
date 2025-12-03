package routes

import (
	"log"
	"net/http"

	"my-go-backend/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/auth/google/login", handlers.GoogleLogin).Methods(http.MethodGet)
	router.HandleFunc("/auth/google/callback", handlers.GoogleCallback).Methods(http.MethodGet)
	router.HandleFunc("/api/me", handlers.Me).Methods(http.MethodGet)
	router.HandleFunc("/api/endpoint", handlers.ApiHandler).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/api/_debug/users", handlers.ListUsers).Methods(http.MethodGet)

	_ = router.Walk(func(route *mux.Route, r *mux.Router, ancestors []*mux.Route) error {
		tmpl, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		log.Printf("registered route: methods=%v path=%s\n", methods, tmpl)
		return nil
	})

	return router
}

