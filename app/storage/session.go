package storage

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions/redis"
	"github.com/mateus-werneck/portifolio/app/tools"
)

func NewSessionStore() redis.Store {
	store, err := redis.NewStore(
        10,
        "tcp",
        os.Getenv("REDIS_HOST"),
        os.Getenv("REDIS_PASSWORD"),
        os.Getenv("REDIS_AUTH"),
    )

	if err != nil {
		tools.GlobalLogger.Error("Failed to create session store", "Error", err)
		log.Fatal(err)
	}

	return store
}
