package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/giwrish/user-service/internal/config"
	"github.com/giwrish/user-service/internal/repository"
	"github.com/go-chi/chi/v5"
)

type UserService struct {
	server *http.Server
	db     *repository.Queries
}

func NewUserService(cfg *config.ServerConfig, queries *repository.Queries) *UserService {

	handler := RegisterRoutes(chi.NewRouter(), queries)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      handler,
		ReadTimeout:  time.Second * time.Duration(cfg.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(cfg.IdleTimeout),
		WriteTimeout: time.Second * time.Duration(cfg.WriteTimeout),
	}

	return &UserService{
		server: server,
		db:     queries,
	}
}

func (svc *UserService) Start() error {
	log.Printf("Starting server on %v", svc.server.Addr)
	return svc.server.ListenAndServe()
}

func (svc *UserService) Shutdown(ctx context.Context) error {
	log.Println("Gracefully shutting down server...")
	return svc.server.Shutdown(ctx)
}
