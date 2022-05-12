// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package env

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgres(connStr string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}
