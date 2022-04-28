package mongo

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type gateway struct {
	logger   common.ILogger
	profiles *mongo.Collection
}

func NewGateway(settings common.ISettings, logger common.ILogger) gateways.IDatabaseGateway {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(settings.DatabaseURI()).SetMaxConnecting(100))

	if err != nil {
		logger.Err(ctx, err, "mongo connection error")
		panic(err)
	}

	db := client.Database(settings.Database())
	profiles := db.Collection("profiles")

	return &gateway{
		logger,
		profiles,
	}
}

type profileBson struct {
	Address           string                  `bson:"_id"`
	DisplayConfig     *displayConfigBson      `bson:"display_config"`
	NonFungibleTokens *[]nonFungibleTokenBson `bson:"non_fungible_tokens"`
	FungibleTokens    *[]fungibleTokenBson    `bson:"fungible_tokens"`
	Statistics        *[]statisticBson        `bson:"statistics"`
	Interactions      *[]interactionBson      `bson:"interactions"`
	Currencies        *[]currencyBson         `json:"currencies"`
}

type contractBson struct {
	Blockchain common.Blockchain `bson:"blockchain"`
	Address    string            `bson:"address"`
	Interface  common.Interface  `bson:"interface"`
}

type nonFungibleTokenBson struct {
	Contract *contractBson `bson:"contract"`
	TokenId  string        `bson:"token_id"`
}

type fungibleTokenBson struct {
	Contract *contractBson `bson:"contract"`
}

type statisticBson struct {
	Contract *contractBson        `bson:"contract"`
	Type     common.StatisticType `bson:"type"`
}

type transactionBson struct {
	Id         string            `bson:"id"`
	Blockchain common.Blockchain `bson:"blockchain"`
}

type interactionBson struct {
	Transaction *transactionBson   `bson:"transaction"`
	Type        common.Interaction `bson:"type"`
	Timestamp   uint64             `bson:"timestamp"`
}

type currencyBson struct {
	Blockchain common.Blockchain `json:"blockchain"`
}

type displayConfigBson struct {
	Colors       *displayColorsBson       `bson:"colors"`
	Info         *displayInfoBson         `bson:"info"`
	Picture      *displayPictureBson      `bson:"picture"`
	Achievements *displayAchievementsBson `bson:"achievements"`
	Groups       *[]displayGroupBson      `bson:"groups"`
}

type displayColorsBson struct {
	Primary       string `bson:"primary"`
	Secondary     string `bson:"secondary"`
	PrimaryText   string `bson:"primary_text"`
	SecondaryText string `bson:"secondary_text"`
	Shadow        string `bson:"shadow"`
	Accent        string `bson:"accent"`
}

type displayInfoBson struct {
	Title         string `bson:"title"`
	Description   string `bson:"description"`
	TwitterHandle string `bson:"twitter_handle"`
}

type displayPictureBson struct {
	Item *displayItemBson `bson:"item,omitempty"` // Item can be nil
}

type displayAchievementsBson struct {
	Text  string                    `bson:"text"`
	Items *[]displayAchievementBson `bson:"items"`
}

type displayAchievementBson struct {
	Id    string                 `bson:"id"`
	Index uint64                 `bson:"index"`
	Type  common.AchievementType `bson:"type"`
}

type displayGroupBson struct {
	Id    string             `bson:"id"`
	Text  string             `bson:"text"`
	Items *[]displayItemBson `bson:"items"`
}

type displayItemBson struct {
	Id    string           `bson:"id"`
	Index uint64           `bson:"index"`
	Type  common.BadgeType `bson:"type"`
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

	currencies := []entities.Currency{}
	for _, currency := range *profileBson.Currencies {
		currencies = append(currencies, entities.Currency{
			Blockchain: currency.Blockchain,
		})
	}

	config := entities.DisplayConfig{
		Colors: &entities.DisplayColors{
			Primary:       profileBson.DisplayConfig.Colors.Primary,
			Secondary:     profileBson.DisplayConfig.Colors.Secondary,
			PrimaryText:   profileBson.DisplayConfig.Colors.PrimaryText,
			SecondaryText: profileBson.DisplayConfig.Colors.SecondaryText,
			Shadow:        profileBson.DisplayConfig.Colors.Shadow,
			Accent:        profileBson.DisplayConfig.Colors.Accent,
		},
		Info: &entities.DisplayInfo{
			Title:         profileBson.DisplayConfig.Info.Title,
			Description:   profileBson.DisplayConfig.Info.Description,
			TwitterHandle: profileBson.DisplayConfig.Info.TwitterHandle,
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

		for _, item := range *groupBson.Items {
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
		Currencies:        &currencies,
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

	currencies := []currencyBson{}
	for _, currency := range *profile.Currencies {
		currencies = append(currencies, currencyBson{
			Blockchain: currency.Blockchain,
		})
	}

	config := displayConfigBson{
		Colors: &displayColorsBson{
			Primary:       profile.DisplayConfig.Colors.Primary,
			Secondary:     profile.DisplayConfig.Colors.Secondary,
			PrimaryText:   profile.DisplayConfig.Colors.PrimaryText,
			SecondaryText: profile.DisplayConfig.Colors.SecondaryText,
			Shadow:        profile.DisplayConfig.Colors.Shadow,
			Accent:        profile.DisplayConfig.Colors.Accent,
		},
		Info: &displayInfoBson{
			Title:         profile.DisplayConfig.Info.Title,
			Description:   profile.DisplayConfig.Info.Description,
			TwitterHandle: profile.DisplayConfig.Info.TwitterHandle,
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
		Currencies:        &currencies,
	}
}
