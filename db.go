package main

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	err := rdb.Conn().Ping(context.Background()).Err()

	if err != nil {
		log.Fatal(err)
	}

	return rdb
}
