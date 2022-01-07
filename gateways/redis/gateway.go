package redis

import (
	"fmt"
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/go-redis/redis/v8"
)

type Gateway struct {
	client *redis.Client
	logger *common.Logger
}

func NewGateway(settings *common.Settings, logger *common.Logger) *Gateway {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", settings.RedisHost, settings.RedisPort),
		Password: settings.RedisPassword,
		DB:       settings.RedisDB,
	})
	return &Gateway{
		client,
		logger,
	}
}

func getFullKey(keys ...string) string {
	return strings.Join(keys, "_")
}
