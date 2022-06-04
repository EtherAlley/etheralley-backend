package entities

import "github.com/etheralley/etheralley-backend/common"

type Contract struct {
	Blockchain common.Blockchain
	Address    string
	Interface  common.Interface
}
