package storage

import (
	"log"
	"log/slog"
	"os"

	"github.com/gin-contrib/sessions/redis"
)

func NewSessionStore() redis.Store {
	store, err := redis.NewStore(10, "tcp", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"), []byte(os.Getenv("REDIS_AUTH")))

	if err != nil {
		slog.Error("Failed to create session store", "Error", err)
		log.Fatal(err)
	}

	return store
}
