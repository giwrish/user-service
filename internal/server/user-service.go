package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/giwrish/user-service/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	server *http.Server
	db     *pgxpool.Pool
}

func NewUserService(cfg *config.ServerConfig, pool *pgxpool.Pool) *UserService {

	handler := RegisterRoutes(chi.NewRouter())

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      handler,
		ReadTimeout:  time.Second * time.Duration(cfg.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(cfg.IdleTimeout),
		WriteTimeout: time.Second * time.Duration(cfg.WriteTimeout),
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

func (svc *UserService) Shutdown(ctx context.Context) error {
	log.Println("Gracefully shutting down server...")
	return svc.server.Shutdown(ctx)
}
