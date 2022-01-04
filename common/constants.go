package common

import "errors"

var ErrNil = errors.New("not found")

const (
	ERC721  string = "ERC721"
	ERC1155 string = "ERC1155"
)

const (
	ETHEREUM string = "ethereum"
)
