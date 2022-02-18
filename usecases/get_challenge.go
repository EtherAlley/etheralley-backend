package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type GetChallengeInput struct {
	Address string `validate:"required,eth_addr"`
}

// Get a challenge message for the provided address
//
// Generate a new challenge and save it to the cache
type IGetChallengeUseCase func(ctx context.Context, input *GetChallengeInput) (*entities.Challenge, error)

func NewGetChallenge(cacheGateway gateways.ICacheGateway) IGetChallengeUseCase {
	return func(ctx context.Context, input *GetChallengeInput) (*entities.Challenge, error) {
		if err := common.ValidateStruct(input); err != nil {
			return nil, err
		}

		challenge := entities.NewChallenge(input.Address)

		err := cacheGateway.SaveChallenge(ctx, challenge)

		return challenge, err
	}
}
