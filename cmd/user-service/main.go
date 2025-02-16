package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/giwrish/user-service/internal/server"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Unable to read config file: %v", err.Error())
	}

	// connect to databse
	pgConnStr := os.Getenv("DATABASE_URL")

	if pgConnStr == "" {
		log.Print("Could not find postgres configuration, using default.")
	}

	dbConfig, err := pgxpool.ParseConfig(pgConnStr)

	if err != nil {
		log.Fatalf("Could not parse database url: %v", err.Error())
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)

	if err != nil {
		log.Fatalf("Could not connect to postgres: %v", err.Error())
	}

	defer dbPool.Close()

	if err = dbPool.Ping(context.Background()); err != nil {
		dbPool.Close()
		log.Fatalf("Could not acquire postgres connection: %v", err.Error())
	}

	svc := server.NewUserService(dbPool)

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
