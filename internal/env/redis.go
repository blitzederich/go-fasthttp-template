// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package env

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func NewRedis(addr, password string, db int) (*redis.Client, error) {
	cache := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	status := cache.Ping(context.Background())
	if err := status.Err(); err != nil {
		return nil, err
	}

	return cache, nil
}
