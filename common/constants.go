package common

import "errors"

var ErrNil = errors.New("not found")

const (
	ERC721   string = "ERC721"
	ERC1155  string = "ERC1155"
	ETHEREUM string = "ethereum"
	POLYGON  string = "polygon"
	ARBITRUM string = "arbitrum"
	OPTIMISM string = "optimism"
)
