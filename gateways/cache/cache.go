package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	redisClient *redis.Client
}

func NewCache() *Cache {
	redisPort := os.Getenv("REDIS_PORT")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%v", redisPort),
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	return &Cache{
		redisClient,
	}
}

func (cache *Cache) Get(ctx context.Context, key string) (string, error) {
	return cache.redisClient.Get(ctx, key).Result()
}

func (cache *Cache) Set(ctx context.Context, key string, value interface{}, exp time.Duration) (string, error) {
	return cache.redisClient.Set(ctx, key, value, exp).Result()
}

func (cache *Cache) Touch(ctx context.Context, key string) error {
	return cache.redisClient.Touch(ctx, key).Err()
}
