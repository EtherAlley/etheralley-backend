package entities

import "github.com/etheralley/etheralley-core-api/common"

type Contract struct {
	Blockchain common.Blockchain
	Address    string
	Interface  common.Interface
}
