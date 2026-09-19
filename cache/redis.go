package cache

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func GetConn() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: 0, 
	})

	defer client.Close()

	return client
}