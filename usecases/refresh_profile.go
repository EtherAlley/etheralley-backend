package usecases

import (
	"context"
	"math/big"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

func NewRefreshProfileUseCase(logger common.ILogger, cacheGateway gateways.ICacheGateway, blockchainGateway gateways.IBlockchainGateway) IRefreshProfileUseCase {
	return &refreshProfileUseCase{
		logger,
		cacheGateway,
		blockchainGateway,
	}
}

type refreshProfileUseCase struct {
	logger            common.ILogger
	cacheGateway      gateways.ICacheGateway
	blockchainGateway gateways.IBlockchainGateway
}

type IRefreshProfileUseCase interface {
	// Refresh any relevant transient info that is currently cached.
	// Info that is currently refreshed:
	//
	// - Store Assets
	Do(ctx context.Context, input *RefreshProfileInput) error
}

type RefreshProfileInput struct {
	Address string `validate:"required,eth_addr"`
}

func (uc *refreshProfileUseCase) Do(ctx context.Context, input *RefreshProfileInput) error {
	if err := common.ValidateStruct(input); err != nil {
		return err
	}

	profile, err := uc.cacheGateway.GetProfileByAddress(ctx, input.Address)

	if err != nil {
		uc.logger.Debug(ctx).Err(err).Msgf("skipping profile refresh %v", input.Address)
		return err
	}

	balances, err := uc.blockchainGateway.GetStoreBalanceBatch(ctx, input.Address, &[]string{common.STORE_PREMIUM, common.STORE_BETA_TESTER})

	if err != nil {
		return err
	}

	profile.StoreAssets.Premium = balances[0].Cmp(big.NewInt(0)) == 1
	profile.StoreAssets.BetaTester = balances[1].Cmp(big.NewInt(0)) == 1

	return uc.cacheGateway.SaveProfile(ctx, profile)
}
