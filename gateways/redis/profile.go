package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/etheralley/etheralley-core-api/entities"
)

const ProfileNamespace = "profile"

func (g *gateway) GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error) {
	profileString, err := g.client.Get(ctx, getFullKey(ProfileNamespace, address)).Result()

	if err != nil {
		return nil, err
	}

	profJson := &profileJson{}
	err = json.Unmarshal([]byte(profileString), profJson)

	if err != nil {
		return nil, err
	}

	profile := fromProfileJson(profJson)

	return profile, nil
}

func (g *gateway) SaveProfile(ctx context.Context, profile *entities.Profile) error {
	profJson := toProfileJson(profile)

	bytes, err := json.Marshal(profJson)

	if err != nil {
		return err
	}

	_, err = g.client.Set(ctx, getFullKey(ProfileNamespace, profile.Address), bytes, time.Hour*24).Result()

	return err
}

func fromProfileJson(profileJson *profileJson) *entities.Profile {
	nfts := []entities.NonFungibleToken{}
	for _, nft := range *profileJson.NonFungibleTokens {
		var metadata *entities.NonFungibleMetadata
		if nft.Metadata != nil {
			metadata = &entities.NonFungibleMetadata{
				Name:        nft.Metadata.Name,
				Description: nft.Metadata.Description,
				Image:       nft.Metadata.Image,
				Attributes:  nft.Metadata.Attributes,
			}
		}
		nfts = append(nfts, entities.NonFungibleToken{
			TokenId: nft.TokenId,
			Contract: &entities.Contract{
				Blockchain: nft.Contract.Blockchain,
				Address:    nft.Contract.Address,
				Interface:  nft.Contract.Interface,
			},
			Balance:  nft.Balance,
			Metadata: metadata,
		})
	}

	tokens := []entities.FungibleToken{}
	for _, token := range *profileJson.FungibleTokens {
		tokens = append(tokens, entities.FungibleToken{
			Contract: &entities.Contract{
				Blockchain: token.Contract.Blockchain,
				Address:    token.Contract.Address,
				Interface:  token.Contract.Interface,
			},
			Balance: token.Balance,
			Metadata: &entities.FungibleMetadata{
				Name:     token.Metadata.Name,
				Symbol:   token.Metadata.Symbol,
				Decimals: token.Metadata.Decimals,
			},
		})
	}

	stats := []entities.Statistic{}
	for _, stat := range *profileJson.Statistics {
		stats = append(stats, entities.Statistic{
			Type: stat.Type,
			Contract: &entities.Contract{
				Blockchain: stat.Contract.Blockchain,
				Address:    stat.Contract.Address,
				Interface:  stat.Contract.Interface,
			},
			Data: stat.Data,
		})
	}

	interactions := []entities.Interaction{}
	for _, interaction := range *profileJson.Interactions {
		interactions = append(interactions, entities.Interaction{
			Type: interaction.Type,
			Transaction: &entities.Transaction{
				Blockchain: interaction.Transaction.Blockchain,
				Id:         interaction.Transaction.Id,
			},
			Timestamp: interaction.Timestamp,
		})
	}

	var config *entities.DisplayConfig
	if profileJson.DisplayConfig != nil {
		config = &entities.DisplayConfig{
			Colors: &entities.DisplayColors{
				Primary:       profileJson.DisplayConfig.Colors.Primary,
				Secondary:     profileJson.DisplayConfig.Colors.Secondary,
				PrimaryText:   profileJson.DisplayConfig.Colors.PrimaryText,
				SecondaryText: profileJson.DisplayConfig.Colors.SecondaryText,
			},
			Text: &entities.DisplayText{
				Title:       profileJson.DisplayConfig.Text.Title,
				Description: profileJson.DisplayConfig.Text.Description,
			},
			Picture: &entities.DisplayPicture{},
			Achievements: &entities.DisplayAchievements{
				Text:  profileJson.DisplayConfig.Achievements.Text,
				Items: &[]entities.DisplayAchievement{},
			},
			Groups: &[]entities.DisplayGroup{},
		}

		if profileJson.DisplayConfig.Picture.Item != nil {
			config.Picture.Item = &entities.DisplayItem{
				Id:    profileJson.DisplayConfig.Picture.Item.Id,
				Index: profileJson.DisplayConfig.Picture.Item.Index,
				Type:  profileJson.DisplayConfig.Picture.Item.Type,
			}
		}

		for _, achievement := range *profileJson.DisplayConfig.Achievements.Items {
			items := append(*config.Achievements.Items, entities.DisplayAchievement{
				Id:    achievement.Id,
				Index: achievement.Index,
				Type:  achievement.Type,
			})
			config.Achievements.Items = &items
		}

		for _, groupJson := range *profileJson.DisplayConfig.Groups {
			group := entities.DisplayGroup{
				Id:    groupJson.Id,
				Text:  groupJson.Text,
				Items: &[]entities.DisplayItem{},
			}

			for _, item := range *groupJson.Items {
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
	}

	return &entities.Profile{
		Address:           profileJson.Address,
		ENSName:           profileJson.ENSName,
		DisplayConfig:     config,
		NonFungibleTokens: &nfts,
		FungibleTokens:    &tokens,
		Statistics:        &stats,
		Interactions:      &interactions,
	}
}

func toProfileJson(profile *entities.Profile) *profileJson {
	nfts := []nonFungibleTokenJson{}
	for _, nft := range *profile.NonFungibleTokens {
		var metadata *nonFungibleMetadataJson
		if nft.Metadata != nil {
			metadata = &nonFungibleMetadataJson{
				Name:        nft.Metadata.Name,
				Description: nft.Metadata.Description,
				Image:       nft.Metadata.Image,
				Attributes:  nft.Metadata.Attributes,
			}
		}
		nfts = append(nfts, nonFungibleTokenJson{
			TokenId: nft.TokenId,
			Contract: &contractJson{
				Blockchain: nft.Contract.Blockchain,
				Address:    nft.Contract.Address,
				Interface:  nft.Contract.Interface,
			},
			Balance:  nft.Balance,
			Metadata: metadata,
		})
	}

	tokens := []fungibleTokenJson{}
	for _, token := range *profile.FungibleTokens {
		tokens = append(tokens, fungibleTokenJson{
			Contract: &contractJson{
				Blockchain: token.Contract.Blockchain,
				Address:    token.Contract.Address,
				Interface:  token.Contract.Interface,
			},
			Balance: token.Balance,
			Metadata: &fungibleMetadataJson{
				Name:     token.Metadata.Name,
				Symbol:   token.Metadata.Symbol,
				Decimals: token.Metadata.Decimals,
			},
		})
	}

	stats := []statisticJson{}
	for _, stat := range *profile.Statistics {
		stats = append(stats, statisticJson{
			Type: stat.Type,
			Contract: &contractJson{
				Blockchain: stat.Contract.Blockchain,
				Address:    stat.Contract.Address,
				Interface:  stat.Contract.Interface,
			},
			Data: stat.Data,
		})
	}

	interactions := []interactionJson{}
	for _, interaction := range *profile.Interactions {
		interactions = append(interactions, interactionJson{
			Type: interaction.Type,
			Transaction: &transactionJson{
				Blockchain: interaction.Transaction.Blockchain,
				Id:         interaction.Transaction.Id,
			},
			Timestamp: interaction.Timestamp,
		})
	}

	var config *displayConfigJson
	if profile.DisplayConfig != nil {
		config = &displayConfigJson{
			Colors: &displayColorsJson{
				Primary:       profile.DisplayConfig.Colors.Primary,
				Secondary:     profile.DisplayConfig.Colors.Secondary,
				PrimaryText:   profile.DisplayConfig.Colors.PrimaryText,
				SecondaryText: profile.DisplayConfig.Colors.SecondaryText,
			},
			Text: &displayTextJson{
				Title:       profile.DisplayConfig.Text.Title,
				Description: profile.DisplayConfig.Text.Description,
			},
			Picture: &displayPictureJson{},
			Achievements: &displayAchievementsJson{
				Text:  profile.DisplayConfig.Achievements.Text,
				Items: &[]displayAchievementJson{},
			},
			Groups: &[]displayGroupJson{},
		}

		if profile.DisplayConfig.Picture.Item != nil {
			config.Picture.Item = &displayItemJson{
				Id:    profile.DisplayConfig.Picture.Item.Id,
				Index: profile.DisplayConfig.Picture.Item.Index,
				Type:  profile.DisplayConfig.Picture.Item.Type,
			}
		}

		for _, achievement := range *profile.DisplayConfig.Achievements.Items {
			items := append(*config.Achievements.Items, displayAchievementJson{
				Id:    achievement.Id,
				Index: achievement.Index,
				Type:  achievement.Type,
			})
			config.Achievements.Items = &items
		}

		for _, group := range *profile.DisplayConfig.Groups {
			groupJson := displayGroupJson{
				Id:    group.Id,
				Text:  group.Text,
				Items: &[]displayItemJson{},
			}

			for _, item := range *group.Items {
				items := append(*groupJson.Items, displayItemJson{
					Id:    item.Id,
					Index: item.Index,
					Type:  item.Type,
				})
				groupJson.Items = &items
			}

			groups := append(*config.Groups, groupJson)
			config.Groups = &groups
		}
	}

	return &profileJson{
		Address:           profile.Address,
		ENSName:           profile.ENSName,
		DisplayConfig:     config,
		NonFungibleTokens: &nfts,
		FungibleTokens:    &tokens,
		Statistics:        &stats,
		Interactions:      &interactions,
	}
}
