package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type SaveProfileInput struct {
	Profile *ProfileInput `validate:"required,dive"`
}

// validate the non transient info being submitted (interactions)
//
// save the profile to the database
//
// remove the cached profile
type ISaveProfileUseCase func(ctx context.Context, input *SaveProfileInput) error

func NewSaveProfile(
	logger common.ILogger,
	databaseGateway gateways.IDatabaseGateway,
	cacheGateway gateways.ICacheGateway,
	getAllInteractions IGetAllInteractionsUseCase,
) ISaveProfileUseCase {
	return func(ctx context.Context, input *SaveProfileInput) error {
		if err := common.ValidateStruct(input); err != nil {
			return err
		}

		// interactions are not transient, so we must validate them on save
		interactions, err := getAllInteractions(ctx, &GetAllInteractionsInput{
			Interactions: toInteractionInputs(input.Profile),
		})

		if err != nil {
			logger.Errf(ctx, err, "failed to validate interactions for profile %v", input.Profile.Address)
			return err
		}

		profile := &entities.Profile{
			Address:           input.Profile.Address,
			Interactions:      interactions,
			NonFungibleTokens: toNonFungibleTokens(input.Profile.NonFungibleTokens),
			FungibleTokens:    toFungibleTokens(input.Profile.FungibleTokens),
			Statistics:        toStatistics(input.Profile.Statistics),
			DisplayConfig:     toDisplayConfig(input.Profile.DisplayConfig),
		}

		err = databaseGateway.SaveProfile(ctx, profile)

		if err != nil {
			logger.Errf(ctx, err, "failed to save profile %v", input.Profile.Address)
			return err
		}

		// users will expect to see there saved changes the next time they visit their profile
		cacheGateway.DeleteProfile(ctx, input.Profile.Address)

		return nil
	}
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
