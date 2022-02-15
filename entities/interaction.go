package entities

import "github.com/etheralley/etheralley-core-api/common"

type Interaction struct {
	Transaction     *Transaction       `bson:"transaction" json:"transaction"`
	Type            common.Interaction `bson:"type" json:"type"`
	TransactionData *TransactionData   `bson:"-" json:"-"`
}
