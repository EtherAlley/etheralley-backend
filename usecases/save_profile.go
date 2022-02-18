package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type SaveProfileInput struct {
	Profile *ProfileInput `validate:"required,dive"`
}

// save the provided profile
//
// fetch transient info for nfts/tokens/stats/etc being submitted
//
// try to save the profile to the cache
//
// regardless of error, save the profile to the database
type ISaveProfileUseCase func(ctx context.Context, input *SaveProfileInput) error

func NewSaveProfile(
	logger common.ILogger,
	cacheGateway gateways.ICacheGateway,
	databaseGateway gateways.IDatabaseGateway,
	getAllNonFungibleTokens IGetAllNonFungibleTokensUseCase,
	getAllFungibleTokens IGetAllFungibleTokensUseCase,
	getAllStatistics IGetAllStatisticsUseCase,
	resolveENSName IResolveENSNameUseCase,
	getAllInteractions IGetAllInteractionsUseCase,
) ISaveProfileUseCase {
	return func(ctx context.Context, input *SaveProfileInput) error {
		if err := common.ValidateStruct(input); err != nil {
			return err
		}

		profile := &entities.Profile{
			Address: input.Profile.Address,
		}

		var wg sync.WaitGroup
		wg.Add(5)

		go func() {
			defer wg.Done()

			// Not all addresses have an ens name. We should not propigate an error for this
			name, _ := resolveENSName(ctx, &ResolveENSNameInput{
				Address: input.Profile.Address,
			})

			profile.ENSName = name
		}()

		go func() {
			defer wg.Done()

			nfts := []GetNonFungibleTokenInput{}
			for _, nft := range *input.Profile.NonFungibleTokens {
				nfts = append(nfts, GetNonFungibleTokenInput{
					Address: input.Profile.Address,
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
			for _, token := range *input.Profile.FungibleTokens {
				tokens = append(tokens, GetFungibleTokenInput{
					Address: input.Profile.Address,
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
			for _, stat := range *input.Profile.Statistics {
				stats = append(stats, GetStatisticsInput{
					Address: input.Profile.Address,
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

			interactions := []GetInteractionInput{}
			for _, interaction := range *input.Profile.Interactions {
				interactions = append(interactions, GetInteractionInput{
					Address: input.Profile.Address,
					Interaction: &InteractionInput{
						Type: interaction.Type,
						Transaction: &TransactionInput{
							Blockchain: interaction.Transaction.Blockchain,
							Id:         interaction.Transaction.Id,
						},
					},
				})
			}

			profile.Interactions = getAllInteractions(ctx, &GetAllInteractionsInput{
				Interactions: &interactions,
			})
		}()

		wg.Wait()

		cacheGateway.SaveProfile(ctx, profile)

		return databaseGateway.SaveProfile(ctx, profile)
	}
}
