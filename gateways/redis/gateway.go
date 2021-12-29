package redis

import (
	"fmt"

	"github.com/eflem00/go-example-app/common"
	"github.com/go-redis/redis/v8"
)

type Gateway struct {
	client *redis.Client
	logger *common.Logger
}

func NewGateway(settings *common.Settings, logger *common.Logger) *Gateway {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%v", settings.RedisPort),
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	return &Gateway{
		client,
		logger,
	}
}

func GetFullKey(namespace string, key string) string {
	return fmt.Sprintf("%v%v", namespace, key)
}
