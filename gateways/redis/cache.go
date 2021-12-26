package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Client *redis.Client
}

func NewCache() *Cache {
	redisPort := os.Getenv("REDIS_PORT")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%v", redisPort),
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	return &Cache{
		Client: redisClient,
	}
}

func GetFullKey(namespace string, key string) string {
	return fmt.Sprintf("%v%v", namespace, key)
}
