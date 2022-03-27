package usecases

import (
	"context"
	"math/big"
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
	blockchainGateway gateways.IBlockchainGateway,
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
			Address:       input.Profile.Address,
			DisplayConfig: toDisplayConfig(input.Profile.DisplayConfig),
		}

		var wg sync.WaitGroup
		wg.Add(6)

		go func() {
			defer wg.Done()
			// Not all addresses have an ens name. We should not propigate an error for this
			profile.ENSName, _ = resolveENSName(ctx, &ResolveENSNameInput{
				Address: input.Profile.Address,
			})
		}()

		go func() {
			defer wg.Done()
			profile.NonFungibleTokens = getAllNonFungibleTokens(ctx, &GetAllNonFungibleTokensInput{
				NonFungibleTokens: toNonFungibleTokenInputs(input.Profile),
			})
		}()

		go func() {
			defer wg.Done()
			profile.FungibleTokens = getAllFungibleTokens(ctx, &GetAllFungibleTokensInput{
				Tokens: toFungibleTokenInputs(input.Profile),
			})
		}()

		go func() {
			defer wg.Done()
			profile.Statistics = getAllStatistics(ctx, &GetAllStatisticsInput{
				Stats: toStatisticInputs(input.Profile),
			})
		}()

		go func() {
			defer wg.Done()
			profile.Interactions = getAllInteractions(ctx, &GetAllInteractionsInput{
				Interactions: toInteractionInputs(input.Profile),
			})
		}()

		go func() {
			defer wg.Done()

			profile.StoreAssets = &entities.StoreAssets{}

			if balances, err := blockchainGateway.GetStoreBalanceBatch(ctx, input.Profile.Address, &[]string{common.STORE_PREMIUM, common.STORE_BETA_TESTER}); err == nil {
				profile.StoreAssets.Premium = balances[0].Cmp(big.NewInt(0)) == 1
				profile.StoreAssets.BetaTester = balances[1].Cmp(big.NewInt(0)) == 1
			}
		}()

		wg.Wait()

		cacheGateway.SaveProfile(ctx, profile)

		return databaseGateway.SaveProfile(ctx, profile)
	}
}

func toNonFungibleTokenInputs(profile *ProfileInput) *[]GetNonFungibleTokenInput {
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
	return &nfts
}

func toFungibleTokenInputs(profile *ProfileInput) *[]GetFungibleTokenInput {
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
	return &tokens
}

func toStatisticInputs(profile *ProfileInput) *[]GetStatisticsInput {
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
	return &stats
}

func toInteractionInputs(profile *ProfileInput) *[]GetInteractionInput {
	interactions := []GetInteractionInput{}
	for _, interaction := range *profile.Interactions {
		interactions = append(interactions, GetInteractionInput{
			Address: profile.Address,
			Interaction: &InteractionInput{
				Type: interaction.Type,
				Transaction: &TransactionInput{
					Blockchain: interaction.Transaction.Blockchain,
					Id:         interaction.Transaction.Id,
				},
			},
		})
	}
	return &interactions
}

func toDisplayConfig(input *DisplayConfigInput) *entities.DisplayConfig {
	config := &entities.DisplayConfig{
		Colors: &entities.DisplayColors{
			Primary:       input.Colors.Primary,
			Secondary:     input.Colors.Secondary,
			PrimaryText:   input.Colors.PrimaryText,
			SecondaryText: input.Colors.SecondaryText,
		},
		Text: &entities.DisplayText{
			Title:       input.Text.Title,
			Description: input.Text.Description,
		},
		Picture: &entities.DisplayPicture{},
		Achievements: &entities.DisplayAchievements{
			Text:  input.Achievements.Text,
			Items: &[]entities.DisplayAchievement{},
		},
		Groups: &[]entities.DisplayGroup{},
	}

	if input.Picture.Item != nil {
		config.Picture.Item = &entities.DisplayItem{
			Id:    input.Picture.Item.Id,
			Index: input.Picture.Item.Index,
			Type:  input.Picture.Item.Type,
		}
	}

	for _, achievement := range *input.Achievements.Items {
		items := append(*config.Achievements.Items, entities.DisplayAchievement{
			Id:    achievement.Id,
			Index: achievement.Index,
			Type:  achievement.Type,
		})
		config.Achievements.Items = &items
	}

	for _, inputGroup := range *input.Groups {
		group := entities.DisplayGroup{
			Id:    inputGroup.Id,
			Text:  inputGroup.Text,
			Items: &[]entities.DisplayItem{},
		}

		for _, inputItem := range *inputGroup.Items {
			items := append(*group.Items, entities.DisplayItem{
				Id:    inputItem.Id,
				Index: inputItem.Index,
				Type:  inputItem.Type,
			})
			group.Items = &items
		}

		groups := append(*config.Groups, group)
		config.Groups = &groups
	}
	return config
}
