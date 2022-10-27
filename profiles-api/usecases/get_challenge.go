package usecases

import (
	"context"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
	"github.com/etheralley/etheralley-backend/profiles-api/gateways"
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

	challenge, err := uc.cacheGateway.GetChallengeByAddress(ctx, input.Address)

	if err == nil {
		return challenge, nil
	}

	challenge = entities.NewChallenge(input.Address)

	err = uc.cacheGateway.SaveChallenge(ctx, challenge)

	return challenge, err
}
