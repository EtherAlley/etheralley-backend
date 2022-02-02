package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

// fetch metadata and ownership of nfts being submitted
// try to save the profile to the cache
// regardless of error, save the profile to the database
func NewSaveProfile(
	logger common.ILogger,
	cacheGateway gateways.ICacheGateway,
	databaseGateway gateways.IDatabaseGateway,
	getAllNonFungibleTokens IGetAllNonFungibleTokensUseCase,
) ISaveProfileUseCase {
	return func(ctx context.Context, profile *entities.Profile) error {
		if err := common.ValidateStruct(profile); err != nil {
			return err
		}

		profile.NonFungibleTokens = getAllNonFungibleTokens(ctx, profile.Address, profile.NonFungibleTokens)

		cacheGateway.SaveProfile(ctx, profile)

		return databaseGateway.SaveProfile(ctx, profile)
	}
}
