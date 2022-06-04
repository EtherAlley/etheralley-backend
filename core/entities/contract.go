package entities

import "github.com/etheralley/etheralley-apis/common"

type Contract struct {
	Blockchain common.Blockchain
	Address    string
	Interface  common.Interface
}
