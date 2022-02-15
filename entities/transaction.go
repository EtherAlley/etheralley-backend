package entities

import "github.com/etheralley/etheralley-core-api/common"

type Transaction struct {
	Id         string            `bson:"id" json:"id" validate:"required"`
	Blockchain common.Blockchain `bson:"blockchain" json:"blockchain" validate:"required,oneof=ethereum polygon arbitrum optimism"`
}

type TransactionData struct {
	From  string  `bson:"-" json:"-"`
	To    *string `bson:"-" json:"-"`
	Data  []byte  `bson:"-" json:"-"`
	Value string  `bson:"-" json:"-"`
}
