package entities

import "github.com/etheralley/etheralley-core-api/common"

type Transaction struct {
	Id         string            `validate:"required"`
	Blockchain common.Blockchain `validate:"required,oneof=ethereum polygon arbitrum optimism"`
}

type TransactionData struct {
	Timestamp uint64
	From      string
	To        *string
	Data      []byte
	Value     string
}
