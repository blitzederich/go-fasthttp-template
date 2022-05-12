// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package env

import (
	"api/internal/config"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	Postgres *pgxpool.Pool
	Redis    *redis.Client
)

func Setup(config *config.Config) error {
	var err error

	Postgres, err = NewPostgres(config.Postgres.ConnStr)
	if err != nil {
		return err
	}

	Redis, err = NewRedis(
		config.Redis.Addr,
		config.Redis.Password,
		config.Redis.Db,
	)
	if err != nil {
		return err
	}

	return nil
}
