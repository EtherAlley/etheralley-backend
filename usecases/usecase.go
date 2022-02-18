package usecases

import (
	"github.com/etheralley/etheralley-core-api/common"
)

type ProfileInput struct {
	Address           string                   `json:"address" validate:"required,eth_addr"`
	ENSName           string                   `json:"ens_name" validate:"required"`
	NonFungibleTokens *[]NonFungibleTokenInput `json:"non_fungible_tokens" validate:"required,dive"`
	FungibleTokens    *[]FungibleTokenInput    `json:"fungible_tokens" validate:"required,dive"`
	Statistics        *[]StatisticInput        `json:"statistics" validate:"required,dive"`
	Interactions      *[]InteractionInput      `json:"interactions" validate:"required,dive"`
}

type ContractInput struct {
	Blockchain common.Blockchain `json:"blockchain" validate:"required,oneof=ethereum polygon arbitrum optimism"`
	Address    string            `json:"address" validate:"required,eth_addr"`
	Interface  common.Interface  `json:"interface" validate:"required,oneof=ERC721 ERC1155 ERC20 ENS_REGISTRAR SUSHISWAP_EXCHANGE UNISWAP_V2_EXCHANGE  UNISWAP_V3_EXCHANGE"`
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
	Type     common.StatisticType `json:"type" validate:"required,oneof=SWAP"`
}

type InteractionInput struct {
	Transaction *TransactionInput  `json:"transaction" validate:"required,dive"`
	Type        common.Interaction `json:"type" validate:"required,oneof=CONTRACT_CREATION SEND_ETHER"`
}
