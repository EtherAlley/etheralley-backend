package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
)

func NewGetChallengeUseCase(cacheGateway *redis.Gateway) GetChallengeUseCase {
	return GetChallenge(cacheGateway)
}

// generate a new challenge and save it to the cache
func GetChallenge(cacheGateway gateways.ICacheGateway) GetChallengeUseCase {
	return func(ctx context.Context, address string) (*entities.Challenge, error) {
		challenge := entities.NewChallenge(address)
		err := cacheGateway.SaveChallenge(ctx, challenge)
		return challenge, err
	}
}
