package redis

import (
	"context"
	"time"

	"github.com/eflem00/go-example-app/common"
)

type AuthCache struct {
	logger *common.Logger
	cache  *Cache
}

func NewAuthCache(cache *Cache, logger *common.Logger) *AuthCache {
	return &AuthCache{
		logger,
		cache,
	}
}

const ChallengeNamespace = "challenge_"

func (ac *AuthCache) GetChallengeMessage(ctx context.Context, address string) (string, error) {
	return ac.cache.Client.Get(ctx, GetFullKey(ChallengeNamespace, address)).Result()
}

func (ac *AuthCache) SetChallengeMessage(ctx context.Context, address string, message string) error {
	_, err := ac.cache.Client.Set(ctx, GetFullKey(ChallengeNamespace, address), message, time.Minute*5).Result()

	return err
}
