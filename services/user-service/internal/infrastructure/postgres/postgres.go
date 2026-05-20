package postgres

import (
	"context"
	"fmt"
	"log"
	"time"
	"users/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnection(cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	var dbName string
	err = pool.QueryRow(ctx, `select current_database()`).Scan(&dbName)
	if err != nil {
		return nil, err
	}

	log.Println("CONNECTED DB:", dbName)

	return pool, nil
}
