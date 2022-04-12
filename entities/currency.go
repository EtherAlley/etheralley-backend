package entities

import "github.com/etheralley/etheralley-core-api/common"

type Currency struct {
	Blockchain common.Blockchain
	Balance    *string // balance can be nil
}
