package usecases

import (
	"context"

	"github.com/etheralley/etheralley-apis/common"
	"github.com/etheralley/etheralley-apis/core/entities"
	"github.com/etheralley/etheralley-apis/core/gateways"
)

func NewGetChallenge(cacheGateway gateways.ICacheGateway) IGetChallengeUseCase {
	return &getChallengeUseCase{
		cacheGateway,
	}
}

type getChallengeUseCase struct {
	cacheGateway gateways.ICacheGateway
}

type IGetChallengeUseCase interface {
	// Get a new challenge message for the provided address
	Do(ctx context.Context, input *GetChallengeInput) (*entities.Challenge, error)
}

type GetChallengeInput struct {
	Address string `validate:"required,eth_addr"`
}

func (uc *getChallengeUseCase) Do(ctx context.Context, input *GetChallengeInput) (*entities.Challenge, error) {
	if err := common.ValidateStruct(input); err != nil {
		return nil, err
	}

	challenge := entities.NewChallenge(input.Address)

	err := uc.cacheGateway.SaveChallenge(ctx, challenge)

	return challenge, err
}
