package entities

import "github.com/etheralley/etheralley-core-api/common"

type Contract struct {
	Blockchain common.Blockchain `bson:"blockchain" json:"blockchain" validate:"required,oneof=ethereum polygon arbitrum optimism"`
	Address    string            `bson:"address" json:"address" validate:"required,eth_addr"`
	Interface  common.Interface  `bson:"interface" json:"interface" validate:"required,oneof=ERC721 ERC1155 ERC20 ENS_REGISTRAR SUSHISWAP_EXCHANGE UNISWAP_V2_EXCHANGE"`
}
