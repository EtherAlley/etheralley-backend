package entities

import "github.com/etheralley/etheralley-backend/common"

type Interaction struct {
	Transaction     *Transaction
	Type            common.Interaction
	Timestamp       uint64
	TransactionData *TransactionData
}
