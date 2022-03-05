package mongo

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (g *gateway) GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error) {
	profileBson := &profileBson{}

	err := g.profiles.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: address}}).Decode(profileBson)

	if err == mongo.ErrNoDocuments {
		return nil, common.ErrNotFound
	}

	profile := fromProfileBson(profileBson)

	return profile, err
}

func (g *gateway) SaveProfile(ctx context.Context, profile *entities.Profile) error {
	profileBson := toProfileBson(profile)

	_, err := g.profiles.UpdateOne(ctx, bson.D{primitive.E{Key: "_id", Value: profile.Address}}, bson.D{primitive.E{Key: "$set", Value: profileBson}}, options.Update().SetUpsert(true))

	return err
}

func fromProfileBson(profileBson *profileBson) *entities.Profile {

	nfts := []entities.NonFungibleToken{}
	for _, nft := range *profileBson.NonFungibleTokens {
		nfts = append(nfts, entities.NonFungibleToken{
			TokenId: nft.TokenId,
			Contract: &entities.Contract{
				Blockchain: nft.Contract.Blockchain,
				Address:    nft.Contract.Address,
				Interface:  nft.Contract.Interface,
			},
		})
	}

	tokens := []entities.FungibleToken{}
	for _, token := range *profileBson.FungibleTokens {
		tokens = append(tokens, entities.FungibleToken{
			Contract: &entities.Contract{
				Blockchain: token.Contract.Blockchain,
				Address:    token.Contract.Address,
				Interface:  token.Contract.Interface,
			},
		})
	}

	stats := []entities.Statistic{}
	for _, stat := range *profileBson.Statistics {
		stats = append(stats, entities.Statistic{
			Type: stat.Type,
			Contract: &entities.Contract{
				Blockchain: stat.Contract.Blockchain,
				Address:    stat.Contract.Address,
				Interface:  stat.Contract.Interface,
			},
		})
	}

	interactions := []entities.Interaction{}
	for _, interaction := range *profileBson.Interactions {
		interactions = append(interactions, entities.Interaction{
			Type: interaction.Type,
			Transaction: &entities.Transaction{
				Blockchain: interaction.Transaction.Blockchain,
				Id:         interaction.Transaction.Id,
			},
			Timestamp: interaction.Timestamp,
		})
	}

	config := entities.DisplayConfig{
		Colors: &entities.DisplayColors{
			Primary:       profileBson.DisplayConfig.Colors.Primary,
			Secondary:     profileBson.DisplayConfig.Colors.Secondary,
			PrimaryText:   profileBson.DisplayConfig.Colors.PrimaryText,
			SecondaryText: profileBson.DisplayConfig.Colors.SecondaryText,
		},
		Text: &entities.DisplayText{
			Title:       profileBson.DisplayConfig.Text.Title,
			Description: profileBson.DisplayConfig.Text.Description,
		},
		Picture: &entities.DisplayPicture{},
		Achievements: &entities.DisplayAchievements{
			Text:  profileBson.DisplayConfig.Achievements.Text,
			Items: &[]entities.DisplayAchievement{},
		},
		Groups: &[]entities.DisplayGroup{},
	}

	if profileBson.DisplayConfig.Picture.Item != nil {
		config.Picture.Item = &entities.DisplayItem{
			Id:    profileBson.DisplayConfig.Picture.Item.Id,
			Index: profileBson.DisplayConfig.Picture.Item.Index,
			Type:  profileBson.DisplayConfig.Picture.Item.Type,
		}
	}

	for _, achievement := range *profileBson.DisplayConfig.Achievements.Items {
		items := append(*config.Achievements.Items, entities.DisplayAchievement{
			Id:    achievement.Id,
			Index: achievement.Index,
			Type:  achievement.Type,
		})
		config.Achievements.Items = &items
	}

	for _, groupBson := range *profileBson.DisplayConfig.Groups {
		group := entities.DisplayGroup{
			Id:    groupBson.Id,
			Text:  groupBson.Text,
			Items: &[]entities.DisplayItem{},
		}

		for _, item := range *group.Items {
			items := append(*group.Items, entities.DisplayItem{
				Id:    item.Id,
				Index: item.Index,
				Type:  item.Type,
			})
			group.Items = &items
		}

		groups := append(*config.Groups, group)
		config.Groups = &groups
	}

	return &entities.Profile{
		Address:           profileBson.Address,
		DisplayConfig:     &config,
		NonFungibleTokens: &nfts,
		FungibleTokens:    &tokens,
		Statistics:        &stats,
		Interactions:      &interactions,
	}
}

func toProfileBson(profile *entities.Profile) *profileBson {
	nfts := []nonFungibleTokenBson{}
	for _, nft := range *profile.NonFungibleTokens {
		nfts = append(nfts, nonFungibleTokenBson{
			TokenId: nft.TokenId,
			Contract: &contractBson{
				Blockchain: nft.Contract.Blockchain,
				Address:    nft.Contract.Address,
				Interface:  nft.Contract.Interface,
			},
		})
	}

	tokens := []fungibleTokenBson{}
	for _, token := range *profile.FungibleTokens {
		tokens = append(tokens, fungibleTokenBson{
			Contract: &contractBson{
				Blockchain: token.Contract.Blockchain,
				Address:    token.Contract.Address,
				Interface:  token.Contract.Interface,
			},
		})
	}

	stats := []statisticBson{}
	for _, stat := range *profile.Statistics {
		stats = append(stats, statisticBson{
			Type: stat.Type,
			Contract: &contractBson{
				Blockchain: stat.Contract.Blockchain,
				Address:    stat.Contract.Address,
				Interface:  stat.Contract.Interface,
			},
		})
	}

	interactions := []interactionBson{}
	for _, interaction := range *profile.Interactions {
		interactions = append(interactions, interactionBson{
			Type: interaction.Type,
			Transaction: &transactionBson{
				Blockchain: interaction.Transaction.Blockchain,
				Id:         interaction.Transaction.Id,
			},
			Timestamp: interaction.Timestamp,
		})
	}

	config := displayConfigBson{
		Colors: &displayColorsBson{
			Primary:       profile.DisplayConfig.Colors.Primary,
			Secondary:     profile.DisplayConfig.Colors.Secondary,
			PrimaryText:   profile.DisplayConfig.Colors.PrimaryText,
			SecondaryText: profile.DisplayConfig.Colors.SecondaryText,
		},
		Text: &displayTextBson{
			Title:       profile.DisplayConfig.Text.Title,
			Description: profile.DisplayConfig.Text.Description,
		},
		Picture: &displayPictureBson{},
		Achievements: &displayAchievementsBson{
			Text:  profile.DisplayConfig.Achievements.Text,
			Items: &[]displayAchievementBson{},
		},
		Groups: &[]displayGroupBson{},
	}

	if profile.DisplayConfig.Picture.Item != nil {
		config.Picture.Item = &displayItemBson{
			Id:    profile.DisplayConfig.Picture.Item.Id,
			Index: profile.DisplayConfig.Picture.Item.Index,
			Type:  profile.DisplayConfig.Picture.Item.Type,
		}
	}

	for _, achievement := range *profile.DisplayConfig.Achievements.Items {
		items := append(*config.Achievements.Items, displayAchievementBson{
			Id:    achievement.Id,
			Index: achievement.Index,
			Type:  achievement.Type,
		})
		config.Achievements.Items = &items
	}

	for _, group := range *profile.DisplayConfig.Groups {
		groupBson := displayGroupBson{
			Id:    group.Id,
			Text:  group.Text,
			Items: &[]displayItemBson{},
		}

		for _, item := range *group.Items {
			items := append(*groupBson.Items, displayItemBson{
				Id:    item.Id,
				Index: item.Index,
				Type:  item.Type,
			})
			groupBson.Items = &items
		}

		groups := append(*config.Groups, groupBson)
		config.Groups = &groups
	}

	return &profileBson{
		Address:           profile.Address,
		DisplayConfig:     &config,
		NonFungibleTokens: &nfts,
		FungibleTokens:    &tokens,
		Statistics:        &stats,
		Interactions:      &interactions,
	}
}
