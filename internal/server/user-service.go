package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	server *http.Server
	db     *pgxpool.Pool
}

func NewUserService(pool *pgxpool.Pool) *UserService {
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		log.Print("Could not find server port, starting on default port 8080")
		port = ":8080"
	}

	handler := RegisterRoutes(chi.NewRouter())

	server := &http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	return &UserService{
		server: server,
		db:     pool,
	}
}

func (svc *UserService) Start() error {
	log.Println("Starting server on :8080")
	return svc.server.ListenAndServe()
}

func (s *UserService) Shutdown(ctx context.Context) error {
	log.Println("Gracefully shutting down server...")
	return s.server.Shutdown(ctx)
}
