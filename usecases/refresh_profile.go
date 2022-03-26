package usecases

import (
	"context"
	"math/big"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type RefreshProfileInput struct {
	Address string `validate:"required,eth_addr"`
}

// Refresh any relevant transient info that is currently cached
//
// Info that is currently refreshed:
//
// - Store Assets
type IRefreshProfileUseCase func(ctx context.Context, input *RefreshProfileInput) error

func NewRefreshProfileUseCase(logger common.ILogger, cacheGateway gateways.ICacheGateway, blockchainGateway gateways.IBlockchainGateway) IRefreshProfileUseCase {
	return func(ctx context.Context, input *RefreshProfileInput) error {
		if err := common.ValidateStruct(input); err != nil {
			return err
		}

		profile, err := cacheGateway.GetProfileByAddress(ctx, input.Address)

		if err != nil {
			return err
		}

		balances, err := blockchainGateway.GetStoreBalanceBatch(ctx, input.Address, &[]string{common.STORE_PREMIUM, common.STORE_BETA_TESTER})

		if err != nil {
			return err
		}

		profile.StoreAssets.Premium = balances[0].Cmp(big.NewInt(0)) == 1
		profile.StoreAssets.BetaTester = balances[1].Cmp(big.NewInt(0)) == 1

		return cacheGateway.SaveProfile(ctx, profile)
	}
}
