package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/giwrish/user-service/internal/server"
)

func main() {
	svc := server.NewUserService()

	// create a channel
	stop := make(chan os.Signal, 1)

	// expect channel of these kinds
	signal.Notify(stop, os.Interrupt, os.Kill)

	// spawn a service and return to main
	go func() {
		if err := svc.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %v", err)
		}
	}()

	// wait for signal
	<-stop
	log.Println("Shutdown signal received. initiating graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := svc.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v", err)
	}

	log.Println("Server gracefully shut down.")

}
