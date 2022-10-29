package redis

import (
	"context"
	"crypto/tls"
	"strings"
	"time"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
	"github.com/etheralley/etheralley-backend/profiles-api/gateways"
	"github.com/etheralley/etheralley-backend/profiles-api/settings"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

type gateway struct {
	settings settings.ISettings
	logger   common.ILogger
	client   *redis.Client
	limiter  *redis_rate.Limiter
}

func NewGateway(settings settings.ISettings, logger common.ILogger) gateways.ICacheGateway {
	return &gateway{
		settings,
		logger,
		nil,
		nil,
	}
}

func (gw *gateway) Init(ctx context.Context) error {
	opt := &redis.Options{
		Addr:      gw.settings.CacheAddr(),
		Password:  gw.settings.CachePassword(),
		DB:        gw.settings.CacheDB(),
		TLSConfig: nil,
	}

	if gw.settings.CacheUseTLS() {
		opt.TLSConfig = &tls.Config{}
	}

	gw.client = redis.NewClient(opt)
	gw.limiter = redis_rate.NewLimiter(gw.client)

	return nil
}

func getFullKey(keys ...string) string {
	return strings.Join(keys, "_")
}

type challengeJson struct {
	Address string    `json:"address"`
	Message string    `json:"message"`
	Expires time.Time `json:"expires"`
}

type profileJson struct {
	Address           string                  `json:"address"`
	Banned            bool                    `json:"banned"`
	LastModified      *time.Time              `json:"last_modified"`
	ENSName           string                  `json:"ens_name"`
	DisplayConfig     *displayConfigJson      `json:"display_config,omitempty"`
	StoreAssets       *storeAssetsJson        `json:"store_assets"`
	ProfilePicture    *nonFungibleTokenJson   `json:"profile_picture"`
	NonFungibleTokens *[]nonFungibleTokenJson `json:"non_fungible_tokens"`
	FungibleTokens    *[]fungibleTokenJson    `json:"fungible_tokens"`
	Statistics        *[]statisticJson        `json:"statistics"`
	Interactions      *[]interactionJson      `json:"interactions"`
	Currencies        *[]currencyJson         `json:"currencies"`
}

type contractJson struct {
	Blockchain common.Blockchain `json:"blockchain"`
	Address    string            `json:"address"`
	Interface  common.Interface  `json:"interface"`
}

type nonFungibleTokenJson struct {
	Contract *contractJson            `json:"contract"`
	TokenId  string                   `json:"token_id"`
	Balance  *string                  `json:"balance"`
	Metadata *nonFungibleMetadataJson `json:"metadata"`
}

type nonFungibleMetadataJson struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Image       string                    `json:"image"`
	Attributes  *[]map[string]interface{} `json:"attributes"`
}

type fungibleTokenJson struct {
	Contract *contractJson         `json:"contract"`
	Balance  *string               `json:"balance"`
	Metadata *fungibleMetadataJson `json:"metadata"`
}

type fungibleMetadataJson struct {
	Name     *string `json:"name"`
	Symbol   *string `json:"symbol"`
	Decimals *uint8  `json:"decimals"`
}

type statisticJson struct {
	Type     common.StatisticType `json:"type"`
	Contract *contractJson        `json:"contract"`
	Data     interface{}          `json:"data"`
}

type interactionJson struct {
	Transaction *transactionJson
	Type        common.Interaction `json:"type"`
	Timestamp   uint64             `json:"timestamp"`
}

type transactionJson struct {
	Id         string            `json:"id"`
	Blockchain common.Blockchain `json:"blockchain"`
}

type currencyJson struct {
	Blockchain common.Blockchain `json:"blockchain"`
	Balance    *string           `json:"balance"`
}

type storeAssetsJson struct {
	Premium    bool `json:"premium"`
	BetaTester bool `json:"beta_tester"`
}

type displayConfigJson struct {
	Colors       *displayColorsJson       `json:"colors"`
	Info         *displayInfoJson         `json:"info"`
	Picture      *displayPictureJson      `json:"picture"`
	Achievements *displayAchievementsJson `json:"achievements"`
	Groups       *[]displayGroupJson      `json:"groups"`
}

type displayColorsJson struct {
	Primary       string `json:"primary"`
	Secondary     string `json:"secondary"`
	PrimaryText   string `json:"primary_text"`
	SecondaryText string `json:"secondary_text"`
	Shadow        string `json:"shadow"`
	Accent        string `json:"accent"`
}

type displayInfoJson struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	TwitterHandle string `json:"twitter_handle"`
}

type displayPictureJson struct {
	Item *displayItemJson `json:"item,omitempty"` // Item can be nil
}

type displayAchievementsJson struct {
	Text  string                    `json:"text"`
	Items *[]displayAchievementJson `json:"items"`
}

type displayAchievementJson struct {
	Id    string                 `json:"id"`
	Index uint64                 `json:"index"`
	Type  common.AchievementType `json:"type"`
}

type displayGroupJson struct {
	Id    string             `json:"id"`
	Text  string             `json:"text"`
	Items *[]displayItemJson `json:"items"`
}

type displayItemJson struct {
	Id    string           `json:"id"`
	Index uint64           `json:"index"`
	Type  common.BadgeType `json:"type"`
}

func fromChallengeJson(challengeJson *challengeJson) *entities.Challenge {
	return &entities.Challenge{
		Address: challengeJson.Address,
		Message: challengeJson.Message,
		Expires: challengeJson.Expires,
	}
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

	currencies := []entities.Currency{}
	for _, currency := range *profileJson.Currencies {
		currencies = append(currencies, entities.Currency{
			Blockchain: currency.Blockchain,
			Balance:    currency.Balance,
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
				Shadow:        profileJson.DisplayConfig.Colors.Shadow,
				Accent:        profileJson.DisplayConfig.Colors.Accent,
			},
			Info: &entities.DisplayInfo{
				Title:         profileJson.DisplayConfig.Info.Title,
				Description:   profileJson.DisplayConfig.Info.Description,
				TwitterHandle: profileJson.DisplayConfig.Info.TwitterHandle,
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

	var profilePicture *entities.NonFungibleToken
	if profileJson.ProfilePicture != nil {
		profilePicture = &entities.NonFungibleToken{
			TokenId: profileJson.ProfilePicture.TokenId,
			Contract: &entities.Contract{
				Blockchain: profileJson.ProfilePicture.Contract.Blockchain,
				Address:    profileJson.ProfilePicture.Contract.Address,
				Interface:  profileJson.ProfilePicture.Contract.Interface,
			},
			Balance: profileJson.ProfilePicture.Balance,
			Metadata: &entities.NonFungibleMetadata{
				Name:        profileJson.ProfilePicture.Metadata.Name,
				Description: profileJson.ProfilePicture.Metadata.Description,
				Image:       profileJson.ProfilePicture.Metadata.Image,
				Attributes:  profileJson.ProfilePicture.Metadata.Attributes,
			},
		}
	}

	return &entities.Profile{
		Address:      profileJson.Address,
		Banned:       profileJson.Banned,
		LastModified: profileJson.LastModified,
		ENSName:      profileJson.ENSName,
		StoreAssets: &entities.StoreAssets{
			Premium:    profileJson.StoreAssets.Premium,
			BetaTester: profileJson.StoreAssets.BetaTester,
		},
		ProfilePicture:    profilePicture,
		DisplayConfig:     config,
		NonFungibleTokens: &nfts,
		FungibleTokens:    &tokens,
		Statistics:        &stats,
		Interactions:      &interactions,
		Currencies:        &currencies,
	}
}

func fromNonFungibleMetadataJson(metadata *nonFungibleMetadataJson) *entities.NonFungibleMetadata {
	return &entities.NonFungibleMetadata{
		Name:        metadata.Name,
		Description: metadata.Description,
		Image:       metadata.Image,
		Attributes:  metadata.Attributes,
	}
}

func fromFungibleMetadataJson(metadata *fungibleMetadataJson) *entities.FungibleMetadata {
	return &entities.FungibleMetadata{
		Name:     metadata.Name,
		Symbol:   metadata.Symbol,
		Decimals: metadata.Decimals,
	}
}

func toChallengeJson(challenge *entities.Challenge) *challengeJson {
	return &challengeJson{
		Address: challenge.Address,
		Message: challenge.Message,
		Expires: challenge.Expires,
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

	currencies := []currencyJson{}
	for _, currency := range *profile.Currencies {
		currencies = append(currencies, currencyJson{
			Blockchain: currency.Blockchain,
			Balance:    currency.Balance,
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
				Shadow:        profile.DisplayConfig.Colors.Shadow,
				Accent:        profile.DisplayConfig.Colors.Accent,
			},
			Info: &displayInfoJson{
				Title:         profile.DisplayConfig.Info.Title,
				Description:   profile.DisplayConfig.Info.Description,
				TwitterHandle: profile.DisplayConfig.Info.TwitterHandle,
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

	var profilePicture *nonFungibleTokenJson
	if profile.ProfilePicture != nil {
		profilePicture = &nonFungibleTokenJson{
			TokenId: profile.ProfilePicture.TokenId,
			Contract: &contractJson{
				Blockchain: profile.ProfilePicture.Contract.Blockchain,
				Address:    profile.ProfilePicture.Contract.Address,
				Interface:  profile.ProfilePicture.Contract.Interface,
			},
			Balance: profile.ProfilePicture.Balance,
			Metadata: &nonFungibleMetadataJson{
				Name:        profile.ProfilePicture.Metadata.Name,
				Description: profile.ProfilePicture.Metadata.Description,
				Image:       profile.ProfilePicture.Metadata.Image,
				Attributes:  profile.ProfilePicture.Metadata.Attributes,
			},
		}
	}

	return &profileJson{
		Address:      profile.Address,
		Banned:       profile.Banned,
		LastModified: profile.LastModified,
		ENSName:      profile.ENSName,
		StoreAssets: &storeAssetsJson{
			Premium:    profile.StoreAssets.Premium,
			BetaTester: profile.StoreAssets.BetaTester,
		},
		ProfilePicture:    profilePicture,
		DisplayConfig:     config,
		NonFungibleTokens: &nfts,
		FungibleTokens:    &tokens,
		Statistics:        &stats,
		Interactions:      &interactions,
		Currencies:        &currencies,
	}
}

func toNonFungibleMetadataJson(metadata *entities.NonFungibleMetadata) *nonFungibleMetadataJson {
	return &nonFungibleMetadataJson{
		Name:        metadata.Name,
		Description: metadata.Description,
		Image:       metadata.Image,
		Attributes:  metadata.Attributes,
	}
}

func toFungibleMetadataJson(metadata *entities.FungibleMetadata) *fungibleMetadataJson {
	return &fungibleMetadataJson{
		Name:     metadata.Name,
		Symbol:   metadata.Symbol,
		Decimals: metadata.Decimals,
	}
}
