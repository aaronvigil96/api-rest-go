package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/postgres"

	pool, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())

	if err != nil {
		return nil, err
	}

	fmt.Println("Conectado a PostgreSQL")

	return pool, nil
}
