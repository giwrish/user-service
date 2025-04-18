package database

import (
	"context"
	"log"
	"time"

	"github.com/giwrish/user-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg *config.DatabaseConfig) *pgxpool.Pool {

	pgConnStr := cfg.Url

	dbConfig, err := pgxpool.ParseConfig(pgConnStr)

	dbConfig.MaxConns = cfg.MaxConnection
	dbConfig.MaxConnIdleTime = time.Duration(cfg.IdleConnections)
	dbConfig.MaxConnLifetime = time.Duration(cfg.MaxConnectionLife)

	if err != nil {
		log.Fatalf("Could not parse database url: %v", err.Error())
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)

	if err != nil {
		log.Fatalf("Could not connect to postgres: %v", err.Error())
	}

	if err = dbPool.Ping(context.Background()); err != nil {
		dbPool.Close()
		log.Fatalf("Could not acquire postgres connection: %v", err.Error())
	}

	return dbPool
}
