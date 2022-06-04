package entities

import "github.com/etheralley/etheralley-apis/common"

type Currency struct {
	Blockchain common.Blockchain
	Balance    *string // balance can be nil
}
