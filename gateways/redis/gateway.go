package redis

import (
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/go-redis/redis/v8"
)

type Gateway struct {
	client *redis.Client
	logger common.ILogger
}

func NewGateway(settings common.ISettings, logger common.ILogger) gateways.ICacheGateway {
	client := redis.NewClient(&redis.Options{
		Addr:     settings.CacheAddr(),
		Password: settings.CachePassword(),
		DB:       settings.CacheDB(),
	})
	return &Gateway{
		client,
		logger,
	}
}

func getFullKey(keys ...string) string {
	return strings.Join(keys, "_")
}

type profileJson struct {
	Address           string                  `json:"address"`
	ENSName           string                  `json:"ens_name"`
	DisplayConfig     *displayConfigJson      `json:"display_config,omitempty"`
	NonFungibleTokens *[]nonFungibleTokenJson `json:"non_fungible_tokens"`
	FungibleTokens    *[]fungibleTokenJson    `json:"fungible_tokens"`
	Statistics        *[]statisticJson        `json:"statistics"`
	Interactions      *[]interactionJson      `json:"interactions"`
}

type contractJson struct {
	Blockchain common.Blockchain `json:"blockchain"`
	Address    string            `json:"address"`
	Interface  common.Interface  `json:"interface"`
}

type nonFungibleTokenJson struct {
	Contract *contractJson            `json:"contract"`
	TokenId  string                   `json:"token_id"`
	Balance  string                   `json:"balance"`
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
	Balance  string                `json:"balance"`
	Metadata *fungibleMetadataJson `json:"metadata"`
}

type fungibleMetadataJson struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
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

type displayConfigJson struct {
	Colors       *displayColorsJson       `json:"colors"`
	Text         *displayTextJson         `json:"text"`
	Picture      *displayPictureJson      `json:"picture"`
	Achievements *displayAchievementsJson `json:"achievements"`
	Groups       *[]displayGroupJson      `json:"groups"`
}

type displayColorsJson struct {
	Primary       string `json:"primary"`
	Secondary     string `json:"secondary"`
	PrimaryText   string `json:"primary_text"`
	SecondaryText string `json:"secondary_text"`
}

type displayTextJson struct {
	Title       string `json:"title"`
	Description string `json:"description"`
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
