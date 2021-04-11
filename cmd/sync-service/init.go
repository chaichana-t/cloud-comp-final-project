package sync_service

import (
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"os"
)

var client *redis.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
}
