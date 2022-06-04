package entities

import "github.com/etheralley/etheralley-backend/common"

type Transaction struct {
	Id         string
	Blockchain common.Blockchain
}

type TransactionData struct {
	Timestamp uint64
	From      string
	To        *string
	Data      []byte
	Value     string
}
