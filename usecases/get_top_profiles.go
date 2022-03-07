package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type GetTopProfilesInput struct {
}

type IGetTopProfilesUseCase func(ctx context.Context, input *GetTopProfilesInput) *[]entities.Profile

func NewGetTopProfilesUseCase(logger common.ILogger, cacheGateway gateways.ICacheGateway) IGetTopProfilesUseCase {
	return func(ctx context.Context, input *GetTopProfilesInput) *[]entities.Profile {
		logger.Info(ctx, "get top profiles usecase")

		cacheGateway.GetTopAddresses(ctx)

		return &[]entities.Profile{}
	}
}
