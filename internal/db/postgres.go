package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool() *pgxpool.Pool {
	dsn := "postgres://postgres:Kashyap@123!@localhost:5432/users_db"

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	return pool
}
