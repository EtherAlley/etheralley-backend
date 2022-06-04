package entities

import "github.com/etheralley/etheralley-backend/common"

type Currency struct {
	Blockchain common.Blockchain
	Balance    *string // balance can be nil
}
