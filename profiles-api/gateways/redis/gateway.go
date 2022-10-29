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
	return &entities.Profile{
		Address:      profileJson.Address,
		Banned:       profileJson.Banned,
		LastModified: profileJson.LastModified,
		ENSName:      profileJson.ENSName,
		StoreAssets: &entities.StoreAssets{
			Premium:    profileJson.StoreAssets.Premium,
			BetaTester: profileJson.StoreAssets.BetaTester,
		},
		ProfilePicture:    fromNonFungibleTokenJson(profileJson.ProfilePicture),
		DisplayConfig:     fromDisplayConfigJson(profileJson.DisplayConfig),
		NonFungibleTokens: fromNonFungibleTokensJson(profileJson.NonFungibleTokens),
		FungibleTokens:    fromFungibleTokensJson(profileJson.FungibleTokens),
		Statistics:        fromStatsJson(profileJson.Statistics),
		Interactions:      fromInteractionsJson(profileJson.Interactions),
		Currencies:        fromCurrenciesJson(profileJson.Currencies),
	}
}

func fromNonFungibleTokensJson(nftsJson *[]nonFungibleTokenJson) *[]entities.NonFungibleToken {
	if nftsJson == nil {
		return nil
	}

	nfts := []entities.NonFungibleToken{}
	for _, nftJson := range *nftsJson {
		nfts = append(nfts, *fromNonFungibleTokenJson(&nftJson))
	}

	return &nfts
}

func fromFungibleTokensJson(tokensJson *[]fungibleTokenJson) *[]entities.FungibleToken {
	if tokensJson == nil {
		return nil
	}

	tokens := []entities.FungibleToken{}
	for _, token := range *tokensJson {
		tokens = append(tokens, entities.FungibleToken{
			Contract: fromContractJson(token.Contract),
			Balance:  token.Balance,
			Metadata: fromFungibleMetadataJson(token.Metadata),
		})
	}

	return &tokens
}

func fromStatsJson(statsJson *[]statisticJson) *[]entities.Statistic {
	if statsJson == nil {
		return nil
	}

	stats := []entities.Statistic{}
	for _, stat := range *statsJson {
		stats = append(stats, entities.Statistic{
			Type:     stat.Type,
			Contract: fromContractJson(stat.Contract),
			Data:     stat.Data,
		})
	}

	return &stats
}

func fromInteractionsJson(interactionsJson *[]interactionJson) *[]entities.Interaction {
	if interactionsJson == nil {
		return nil
	}

	interactions := []entities.Interaction{}
	for _, interaction := range *interactionsJson {
		interactions = append(interactions, entities.Interaction{
			Type: interaction.Type,
			Transaction: &entities.Transaction{
				Blockchain: interaction.Transaction.Blockchain,
				Id:         interaction.Transaction.Id,
			},
			Timestamp: interaction.Timestamp,
		})
	}

	return &interactions
}

func fromCurrenciesJson(currenciesJson *[]currencyJson) *[]entities.Currency {
	if currenciesJson == nil {
		return nil
	}

	currencies := []entities.Currency{}
	for _, currency := range *currenciesJson {
		currencies = append(currencies, entities.Currency{
			Blockchain: currency.Blockchain,
			Balance:    currency.Balance,
		})
	}

	return &currencies
}

func fromDisplayConfigJson(configJson *displayConfigJson) *entities.DisplayConfig {
	if configJson == nil {
		return nil
	}

	config := &entities.DisplayConfig{
		Colors: &entities.DisplayColors{
			Primary:       configJson.Colors.Primary,
			Secondary:     configJson.Colors.Secondary,
			PrimaryText:   configJson.Colors.PrimaryText,
			SecondaryText: configJson.Colors.SecondaryText,
			Shadow:        configJson.Colors.Shadow,
			Accent:        configJson.Colors.Accent,
		},
		Info: &entities.DisplayInfo{
			Title:         configJson.Info.Title,
			Description:   configJson.Info.Description,
			TwitterHandle: configJson.Info.TwitterHandle,
		},
		Achievements: &entities.DisplayAchievements{
			Text:  configJson.Achievements.Text,
			Items: &[]entities.DisplayAchievement{},
		},
		Groups: &[]entities.DisplayGroup{},
	}

	for _, achievement := range *configJson.Achievements.Items {
		items := append(*config.Achievements.Items, entities.DisplayAchievement{
			Id:    achievement.Id,
			Index: achievement.Index,
			Type:  achievement.Type,
		})
		config.Achievements.Items = &items
	}

	for _, groupJson := range *configJson.Groups {
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

	return config
}

func fromNonFungibleTokenJson(nft *nonFungibleTokenJson) *entities.NonFungibleToken {
	if nft == nil {
		return nil
	}

	return &entities.NonFungibleToken{
		TokenId:  nft.TokenId,
		Contract: fromContractJson(nft.Contract),
		Balance:  nft.Balance,
		Metadata: fromNonFungibleMetadataJson(nft.Metadata),
	}
}

func fromContractJson(contract *contractJson) *entities.Contract {
	return &entities.Contract{
		Blockchain: contract.Blockchain,
		Address:    contract.Address,
		Interface:  contract.Interface,
	}
}

func fromNonFungibleMetadataJson(metadata *nonFungibleMetadataJson) *entities.NonFungibleMetadata {
	if metadata == nil {
		return nil
	}

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
	return &profileJson{
		Address:      profile.Address,
		Banned:       profile.Banned,
		LastModified: profile.LastModified,
		ENSName:      profile.ENSName,
		StoreAssets: &storeAssetsJson{
			Premium:    profile.StoreAssets.Premium,
			BetaTester: profile.StoreAssets.BetaTester,
		},
		ProfilePicture:    toNonFungibleTokenJson(profile.ProfilePicture),
		DisplayConfig:     toDisplayConfigJson(profile.DisplayConfig),
		NonFungibleTokens: toNonFungibleTokensJson(profile.NonFungibleTokens),
		FungibleTokens:    toFungibleTokensJson(profile.FungibleTokens),
		Statistics:        toStatsJson(profile.Statistics),
		Interactions:      toInteractionsJson(profile.Interactions),
		Currencies:        toCurrenciesJson(profile.Currencies),
	}
}

func toNonFungibleTokensJson(nfts *[]entities.NonFungibleToken) *[]nonFungibleTokenJson {
	if nfts == nil {
		return nil
	}

	nftsJson := []nonFungibleTokenJson{}
	for _, nft := range *nfts {
		nftsJson = append(nftsJson, *toNonFungibleTokenJson(&nft))
	}

	return &nftsJson
}

func toFungibleTokensJson(tokens *[]entities.FungibleToken) *[]fungibleTokenJson {
	if tokens == nil {
		return nil
	}

	tokensJson := []fungibleTokenJson{}
	for _, token := range *tokens {
		tokensJson = append(tokensJson, fungibleTokenJson{
			Contract: toContractJosn(token.Contract),
			Balance:  token.Balance,
			Metadata: toFungibleMetadataJson(token.Metadata),
		})
	}

	return &tokensJson
}

func toStatsJson(stats *[]entities.Statistic) *[]statisticJson {
	if stats == nil {
		return nil
	}

	statsJson := []statisticJson{}
	for _, stat := range *stats {
		statsJson = append(statsJson, statisticJson{
			Type:     stat.Type,
			Contract: toContractJosn(stat.Contract),
			Data:     stat.Data,
		})
	}

	return &statsJson
}

func toInteractionsJson(interactions *[]entities.Interaction) *[]interactionJson {
	if interactions == nil {
		return nil
	}

	interactionsJson := []interactionJson{}
	for _, interaction := range *interactions {
		interactionsJson = append(interactionsJson, interactionJson{
			Type: interaction.Type,
			Transaction: &transactionJson{
				Blockchain: interaction.Transaction.Blockchain,
				Id:         interaction.Transaction.Id,
			},
			Timestamp: interaction.Timestamp,
		})
	}

	return &interactionsJson
}

func toCurrenciesJson(currencies *[]entities.Currency) *[]currencyJson {
	if currencies == nil {
		return nil
	}

	currenciesJson := []currencyJson{}
	for _, currency := range *currencies {
		currenciesJson = append(currenciesJson, currencyJson{
			Blockchain: currency.Blockchain,
			Balance:    currency.Balance,
		})
	}

	return &currenciesJson
}

func toDisplayConfigJson(config *entities.DisplayConfig) *displayConfigJson {
	if config == nil {
		return nil
	}

	configJson := &displayConfigJson{
		Colors: &displayColorsJson{
			Primary:       config.Colors.Primary,
			Secondary:     config.Colors.Secondary,
			PrimaryText:   config.Colors.PrimaryText,
			SecondaryText: config.Colors.SecondaryText,
			Shadow:        config.Colors.Shadow,
			Accent:        config.Colors.Accent,
		},
		Info: &displayInfoJson{
			Title:         config.Info.Title,
			Description:   config.Info.Description,
			TwitterHandle: config.Info.TwitterHandle,
		},
		Achievements: &displayAchievementsJson{
			Text:  config.Achievements.Text,
			Items: &[]displayAchievementJson{},
		},
		Groups: &[]displayGroupJson{},
	}

	for _, achievement := range *config.Achievements.Items {
		items := append(*configJson.Achievements.Items, displayAchievementJson{
			Id:    achievement.Id,
			Index: achievement.Index,
			Type:  achievement.Type,
		})
		configJson.Achievements.Items = &items
	}

	for _, group := range *config.Groups {
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

		groups := append(*configJson.Groups, groupJson)
		configJson.Groups = &groups
	}

	return configJson
}

func toNonFungibleTokenJson(nft *entities.NonFungibleToken) *nonFungibleTokenJson {
	if nft == nil {
		return nil
	}

	return &nonFungibleTokenJson{
		TokenId:  nft.TokenId,
		Contract: toContractJosn(nft.Contract),
		Balance:  nft.Balance,
		Metadata: toNonFungibleMetadataJson(nft.Metadata),
	}
}

func toContractJosn(contract *entities.Contract) *contractJson {
	return &contractJson{
		Blockchain: contract.Blockchain,
		Address:    contract.Address,
		Interface:  contract.Interface,
	}
}

func toNonFungibleMetadataJson(metadata *entities.NonFungibleMetadata) *nonFungibleMetadataJson {
	if metadata == nil {
		return nil
	}

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
