package redis

import (
	"context"
	"time"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/entities"
)

type ChallengeCache struct {
	logger *common.Logger
	cache  *Cache
}

func NewChallengeCache(cache *Cache, logger *common.Logger) *ChallengeCache {
	return &ChallengeCache{
		logger,
		cache,
	}
}

const ChallengeNamespace = "challenge_"

func (cc *ChallengeCache) GetChallenge(ctx context.Context, address string) (*entities.Challenge, error) {
	msg, err := cc.cache.Client.Get(ctx, GetFullKey(ChallengeNamespace, address)).Result()
	return &entities.Challenge{Message: msg}, err
}

func (cc *ChallengeCache) SetChallenge(ctx context.Context, address string, challenge *entities.Challenge) error {
	_, err := cc.cache.Client.Set(ctx, GetFullKey(ChallengeNamespace, address), challenge.Message, time.Minute*5).Result()

	return err
}
