package server

import (
	"net/http"

	"github.com/giwrish/user-service/internal/handlers"
	"github.com/giwrish/user-service/internal/repository"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router *chi.Mux, queries *repository.Queries) http.Handler {

	users := handlers.NewUserHandler(queries)

	router.Route("/api/user", func(router chi.Router) {
		router.Post("/", users.CreateUser)
		router.Get("/{username}", users.GetUser)
	})
	return router
}
