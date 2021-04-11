package sync_service

import (
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"os"
)

var client *redis.Client

var cache map[string]int

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	cache = make(map[string]int)
}
