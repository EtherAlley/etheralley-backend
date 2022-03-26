package usecases

import (
	"context"
	"math/big"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type GetProfileInput struct {
	Address string `validate:"required,eth_addr"`
}

// get the profile for the provided address
type IGetProfileUseCase func(ctx context.Context, input *GetProfileInput) (*entities.Profile, error)

// first try to get the profile from the cache.
// if cache miss, go to database
// if database miss, fetch default profile
// if database hit, re-fetch transient info
func NewGetProfile(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	cacheGateway gateways.ICacheGateway,
	databaseGateway gateways.IDatabaseGateway,
	getDefaultProfile IGetDefaultProfileUseCase,
	getAllNonFungibleTokens IGetAllNonFungibleTokensUseCase,
	getAllFungibleTokens IGetAllFungibleTokensUseCase,
	getAllStatistics IGetAllStatisticsUseCase,
	resolveENSName IResolveENSNameUseCase,
) IGetProfileUseCase {
	return func(ctx context.Context, input *GetProfileInput) (*entities.Profile, error) {
		if err := common.ValidateStruct(input); err != nil {
			return nil, err
		}

		profile, err := cacheGateway.GetProfileByAddress(ctx, input.Address)

		if err == nil {
			logger.Debugf(ctx, "cache hit for profile %v", input.Address)
			return profile, nil
		}

		logger.Debugf(ctx, "cache miss for profile %v", input.Address)

		profile, err = databaseGateway.GetProfileByAddress(ctx, input.Address)

		if err == common.ErrNotFound {
			logger.Debugf(ctx, "db miss for profile %v", input.Address)

			profile, err := getDefaultProfile(ctx, &GetDefaultProfileInput{
				Address: input.Address,
			})

			if err != nil {
				logger.Err(ctx, err, "err getting default profile")
				return nil, err
			}

			cacheGateway.SaveProfile(ctx, profile)

			return profile, nil
		}

		if err != nil {
			logger.Err(ctx, err, "err getting profile from db")
			return nil, err
		}

		logger.Debugf(ctx, "db hit for profile %v", input.Address)

		var wg sync.WaitGroup
		wg.Add(5)

		go func() {
			defer wg.Done()

			// Not all addresses have an ens name. We should not propigate an error for this
			name, _ := resolveENSName(ctx, &ResolveENSNameInput{
				Address: profile.Address,
			})

			profile.ENSName = name
		}()

		go func() {
			defer wg.Done()

			nfts := []GetNonFungibleTokenInput{}
			for _, nft := range *profile.NonFungibleTokens {
				nfts = append(nfts, GetNonFungibleTokenInput{
					Address: profile.Address,
					NonFungibleToken: &NonFungibleTokenInput{
						TokenId: nft.TokenId,
						Contract: &ContractInput{
							Blockchain: nft.Contract.Blockchain,
							Address:    nft.Contract.Address,
							Interface:  nft.Contract.Interface,
						},
					},
				})
			}

			profile.NonFungibleTokens = getAllNonFungibleTokens(ctx, &GetAllNonFungibleTokensInput{
				NonFungibleTokens: &nfts,
			})
		}()

		go func() {
			defer wg.Done()

			tokens := []GetFungibleTokenInput{}
			for _, token := range *profile.FungibleTokens {
				tokens = append(tokens, GetFungibleTokenInput{
					Address: profile.Address,
					Token: &FungibleTokenInput{
						Contract: &ContractInput{
							Blockchain: token.Contract.Blockchain,
							Address:    token.Contract.Address,
							Interface:  token.Contract.Interface,
						},
					},
				})
			}

			profile.FungibleTokens = getAllFungibleTokens(ctx, &GetAllFungibleTokensInput{
				Tokens: &tokens,
			})
		}()

		go func() {
			defer wg.Done()

			stats := []GetStatisticsInput{}
			for _, stat := range *profile.Statistics {
				stats = append(stats, GetStatisticsInput{
					Address: profile.Address,
					Statistic: &StatisticInput{
						Type: stat.Type,
						Contract: &ContractInput{
							Blockchain: stat.Contract.Blockchain,
							Address:    stat.Contract.Address,
							Interface:  stat.Contract.Interface,
						},
					},
				})
			}

			profile.Statistics = getAllStatistics(ctx, &GetAllStatisticsInput{
				Stats: &stats,
			})
		}()

		go func() {
			defer wg.Done()

			profile.StoreAssets = &entities.StoreAssets{}

			if balances, err := blockchainGateway.GetStoreBalanceBatch(ctx, input.Address, &[]string{common.STORE_PREMIUM, common.STORE_BETA_TESTER}); err == nil {
				profile.StoreAssets.Premium = balances[0].Cmp(big.NewInt(0)) == 1
				profile.StoreAssets.BetaTester = balances[1].Cmp(big.NewInt(0)) == 1
			}
		}()

		wg.Wait()

		cacheGateway.SaveProfile(ctx, profile)

		return profile, nil
	}
}
