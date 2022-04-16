package usecases

import (
	"github.com/etheralley/etheralley-core-api/common"
)

type ProfileInput struct {
	Address           string                   `json:"-" validate:"required,eth_addr"`
	DisplayConfig     *DisplayConfigInput      `json:"display_config" validate:"required,dive"`
	NonFungibleTokens *[]NonFungibleTokenInput `json:"non_fungible_tokens" validate:"required,dive"`
	FungibleTokens    *[]FungibleTokenInput    `json:"fungible_tokens" validate:"required,dive"`
	Statistics        *[]StatisticInput        `json:"statistics" validate:"required,dive"`
	Interactions      *[]InteractionInput      `json:"interactions" validate:"required,dive"`
	Currencies        *[]CurrencyInput         `json:"currencies" validate:"required,dive"`
}

type ContractInput struct {
	Blockchain common.Blockchain `json:"blockchain" validate:"required,oneof=ethereum polygon arbitrum optimism"`
	Address    string            `json:"address" validate:"required,eth_addr"`
	Interface  common.Interface  `json:"interface" validate:"required,oneof=ERC721 ERC1155 ERC20 ENS_REGISTRAR SUSHISWAP_EXCHANGE UNISWAP_V2_EXCHANGE UNISWAP_V3_EXCHANGE ROCKET_POOL"`
}

type TransactionInput struct {
	Id         string            `json:"id" validate:"required"`
	Blockchain common.Blockchain `json:"blockchain" validate:"required,oneof=ethereum polygon arbitrum optimism"`
}

type NonFungibleTokenInput struct {
	Contract *ContractInput `json:"contract" validate:"required,dive"`
	TokenId  string         `json:"token_id" validate:"required,numeric"`
}

type FungibleTokenInput struct {
	Contract *ContractInput `json:"contract" validate:"required,dive"`
}

type StatisticInput struct {
	Contract *ContractInput       `json:"contract" validate:"required,dive"`
	Type     common.StatisticType `json:"type" validate:"required,oneof=SWAP STAKE"`
}

type InteractionInput struct {
	Transaction *TransactionInput  `json:"transaction" validate:"required,dive"`
	Type        common.Interaction `json:"type" validate:"required,oneof=CONTRACT_CREATION SEND_ETHER"`
}

type CurrencyInput struct {
	Blockchain common.Blockchain `json:"blockchain" validate:"required,oneof=ethereum polygon arbitrum optimism"`
}

type DisplayConfigInput struct {
	Colors       *DisplayColorsInput       `json:"colors" validate:"required,dive"`
	Info         *DisplayInfoInput         `json:"info" validate:"required,dive"`
	Picture      *DisplayPictureInput      `json:"picture" validate:"required,dive"`
	Achievements *DisplayAchievementsInput `json:"achievements" validate:"required,dive"`
	Groups       *[]DisplayGroupInput      `json:"groups" validate:"required,dive"`
}

type DisplayColorsInput struct {
	Primary       string `json:"primary" validate:"required,max=15"`
	Secondary     string `json:"secondary" validate:"required,max=15"`
	PrimaryText   string `json:"primary_text" validate:"required,max=15"`
	SecondaryText string `json:"secondary_text" validate:"required,max=15"`
}

type DisplayInfoInput struct {
	Title         string `json:"title" validate:"max=40"`
	Description   string `json:"description" validate:"max=500"`
	TwitterHandle string `json:"twitter_handle" validate:"max=15"`
}

type DisplayPictureInput struct {
	Item *DisplayItemInput `json:"item,omitempty" validate:""` // Item can be nil. TODO: figure out how to validate this properly. dive will not allow nil values
}

type DisplayAchievementsInput struct {
	Text  string                     `json:"text" validate:"max=30"`
	Items *[]DisplayAchievementInput `json:"items" validate:"required,dive"`
}

type DisplayAchievementInput struct {
	Id    string                 `json:"id" validate:"required"`
	Index uint64                 `json:"index" validate:"gte=0"`
	Type  common.AchievementType `json:"type" validate:"required,oneof=interactions"`
}

type DisplayGroupInput struct {
	Id    string              `json:"id" validate:"required,max=30"`
	Text  string              `json:"text" validate:"max=30"`
	Items *[]DisplayItemInput `json:"items" validate:"required,dive"`
}

type DisplayItemInput struct {
	Id    string           `json:"id" validate:"required"`
	Index uint64           `json:"index" validate:"gte=0"`
	Type  common.BadgeType `json:"type" validate:"required,oneof=non_fungible_tokens fungible_tokens statistics currencies"`
}
