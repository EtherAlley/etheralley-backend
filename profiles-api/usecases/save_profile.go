package usecases

import (
	"context"
	"fmt"
	"math/big"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
	"github.com/etheralley/etheralley-backend/profiles-api/gateways"
)

func NewSaveProfile(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	databaseGateway gateways.IDatabaseGateway,
	cacheGateway gateways.ICacheGateway,
	getAllInteractions IGetAllInteractionsUseCase,
) ISaveProfileUseCase {
	return &saveProfileUseCase{
		logger,
		blockchainGateway,
		databaseGateway,
		cacheGateway,
		getAllInteractions,
	}
}

type saveProfileUseCase struct {
	logger             common.ILogger
	blockchainGateway  gateways.IBlockchainGateway
	databaseGateway    gateways.IDatabaseGateway
	cacheGateway       gateways.ICacheGateway
	getAllInteractions IGetAllInteractionsUseCase
}

type ISaveProfileUseCase interface {
	// Validate the non transient info being submitted (interactions).
	// Save the profile to the database.
	// Remove the cached profile.
	//
	// TODO: can fetch interactions and premium balance concurrently.
	Do(ctx context.Context, input *SaveProfileInput) error
}

type SaveProfileInput struct {
	Profile *ProfileInput `validate:"required,dive"`
}

func (uc *saveProfileUseCase) Do(ctx context.Context, input *SaveProfileInput) error {
	// general validation on profile being submitted
	if err := common.ValidateStruct(input); err != nil {
		return err
	}

	// validate total badge count
	badgeCount := len(*input.Profile.FungibleTokens) + len(*input.Profile.NonFungibleTokens) + len(*input.Profile.Statistics) + len(*input.Profile.Currencies)
	balances, err := uc.blockchainGateway.GetStoreBalanceBatch(ctx, input.Profile.Address, &[]string{common.STORE_PREMIUM})

	if err != nil {
		return fmt.Errorf("failed to get premium balance %w", err)
	}

	if balances[0].Cmp(big.NewInt(0)) == 1 && badgeCount > int(common.PREMIUM_TOTAL_BADGE_COUNT) {
		return fmt.Errorf("invalid total badge count for premium provided %v %v", balances[0], badgeCount)
	} else if balances[0].Cmp(big.NewInt(0)) == 0 && badgeCount > int(common.REGULAR_TOTAL_BADGE_COUNT) {
		return fmt.Errorf("invalid total badge count for regular provided %v %v", balances[0], badgeCount)
	}

	// validate that all interactions being submitted are unique
	typeMap := map[string]bool{}
	for _, interaction := range *input.Profile.Interactions {
		if typeMap[interaction.Type] {
			return fmt.Errorf("duplicate interaction type detected %v", interaction.Type)
		}
		typeMap[interaction.Type] = true
	}

	// interactions are not transient, so we must validate them against the on-chain transaction on every save
	interactions, err := uc.getAllInteractions.Do(ctx, &GetAllInteractionsInput{
		Interactions: toInteractionInputs(input.Profile),
	})

	if err != nil {
		return fmt.Errorf("failed to validate interactions %w", err)
	}

	profile := &entities.Profile{
		Address:           input.Profile.Address,
		Interactions:      interactions,
		NonFungibleTokens: toNonFungibleTokens(input.Profile.NonFungibleTokens),
		FungibleTokens:    toFungibleTokens(input.Profile.FungibleTokens),
		Statistics:        toStatistics(input.Profile.Statistics),
		Currencies:        toCurrencies(input.Profile.Currencies),
		DisplayConfig:     toDisplayConfig(input.Profile.DisplayConfig),
	}

	err = uc.databaseGateway.SaveProfile(ctx, profile)

	if err != nil {
		return fmt.Errorf("failed to save profile to db %w", err)
	}

	// users will expect to see there saved changes the next time they visit their profile
	uc.cacheGateway.DeleteProfile(ctx, input.Profile.Address)

	return nil
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

func toNonFungibleTokens(nftInputs *[]NonFungibleTokenInput) *[]entities.NonFungibleToken {
	nfts := []entities.NonFungibleToken{}
	for _, nft := range *nftInputs {
		nfts = append(nfts, entities.NonFungibleToken{
			TokenId: nft.TokenId,
			Contract: &entities.Contract{
				Blockchain: nft.Contract.Blockchain,
				Address:    nft.Contract.Address,
				Interface:  nft.Contract.Interface,
			},
		})
	}
	return &nfts
}

func toFungibleTokens(tokenInputs *[]FungibleTokenInput) *[]entities.FungibleToken {
	tokens := []entities.FungibleToken{}
	for _, token := range *tokenInputs {
		tokens = append(tokens, entities.FungibleToken{
			Contract: &entities.Contract{
				Blockchain: token.Contract.Blockchain,
				Address:    token.Contract.Address,
				Interface:  token.Contract.Interface,
			},
		})
	}
	return &tokens
}

func toStatistics(statisticInputs *[]StatisticInput) *[]entities.Statistic {
	stats := []entities.Statistic{}
	for _, stat := range *statisticInputs {
		stats = append(stats, entities.Statistic{
			Type: stat.Type,
			Contract: &entities.Contract{
				Blockchain: stat.Contract.Blockchain,
				Address:    stat.Contract.Address,
				Interface:  stat.Contract.Interface,
			},
		})
	}
	return &stats
}

func toCurrencies(currencyInputs *[]CurrencyInput) *[]entities.Currency {
	currencies := []entities.Currency{}
	for _, currency := range *currencyInputs {
		currencies = append(currencies, entities.Currency{
			Blockchain: currency.Blockchain,
		})
	}
	return &currencies
}

func toDisplayConfig(input *DisplayConfigInput) *entities.DisplayConfig {
	config := &entities.DisplayConfig{
		Colors: &entities.DisplayColors{
			Primary:       input.Colors.Primary,
			Secondary:     input.Colors.Secondary,
			PrimaryText:   input.Colors.PrimaryText,
			SecondaryText: input.Colors.SecondaryText,
			Shadow:        input.Colors.Shadow,
			Accent:        input.Colors.Accent,
		},
		Info: &entities.DisplayInfo{
			Title:         input.Info.Title,
			Description:   input.Info.Description,
			TwitterHandle: input.Info.TwitterHandle,
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
