package usecases

import (
	"context"

	"github.com/eflem00/go-example-app/entities"
	"github.com/eflem00/go-example-app/gateways"
	"github.com/eflem00/go-example-app/gateways/redis"
)

type IGetChallengeUseCase interface {
	Go(ctx context.Context, address string) (*entities.Challenge, error)
}

type GetChallengeUseCase struct {
	cacheGateway gateways.ICacheGateway
}

func NewGetChallengeUseCase(cacheGateway *redis.Gateway) *GetChallengeUseCase {
	return &GetChallengeUseCase{
		cacheGateway,
	}
}

func (uc *GetChallengeUseCase) Go(ctx context.Context, address string) (*entities.Challenge, error) {
	challenge := entities.NewChallenge(address)
	err := uc.cacheGateway.SaveChallenge(ctx, challenge)
	return challenge, err
}
