package redis

import (
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/go-redis/redis/v8"
)

type Gateway struct {
	client *redis.Client
	logger common.ILogger
}

func NewGateway(settings common.ISettings, logger common.ILogger) gateways.ICacheGateway {
	client := redis.NewClient(&redis.Options{
		Addr:     settings.CacheAddr(),
		Password: settings.CachePassword(),
		DB:       settings.CacheDB(),
	})
	return &Gateway{
		client,
		logger,
	}
}

func getFullKey(keys ...string) string {
	return strings.Join(keys, "_")
}
