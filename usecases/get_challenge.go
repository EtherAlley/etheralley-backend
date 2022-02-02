package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

// generate a new challenge and save it to the cache
func NewGetChallenge(cacheGateway gateways.ICacheGateway) IGetChallengeUseCase {
	return func(ctx context.Context, address string) (*entities.Challenge, error) {
		if err := common.ValidateField(address, `required,eth_addr`); err != nil {
			return nil, err
		}

		challenge := entities.NewChallenge(address)

		err := cacheGateway.SaveChallenge(ctx, challenge)

		return challenge, err
	}
}
