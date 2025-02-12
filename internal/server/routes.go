package server

import (
	"net/http"

	"github.com/giwrish/user-service/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router *chi.Mux) http.Handler {
	router.Route("/api/user", func(router chi.Router) {
		router.Post("/", handlers.CreateUser)
		router.Get("/{username}", handlers.GetUser)
	})
	return router
}
