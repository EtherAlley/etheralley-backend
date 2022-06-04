package entities

import "github.com/etheralley/etheralley-core-api/common"

type Interaction struct {
	Transaction     *Transaction
	Type            common.Interaction
	Timestamp       uint64
	TransactionData *TransactionData
}
