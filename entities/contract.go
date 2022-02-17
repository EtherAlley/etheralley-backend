package entities

import "github.com/etheralley/etheralley-core-api/common"

type Contract struct {
	Blockchain common.Blockchain `validate:"required,oneof=ethereum polygon arbitrum optimism"`
	Address    string            `validate:"required,eth_addr"`
	Interface  common.Interface  `validate:"required,oneof=ERC721 ERC1155 ERC20 ENS_REGISTRAR SUSHISWAP_EXCHANGE UNISWAP_V2_EXCHANGE  UNISWAP_V3_EXCHANGE"`
}
