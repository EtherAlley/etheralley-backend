package entities

type FungibleToken struct {
	Contract *Contract         `bson:"contract" json:"contract" validate:"required"`
	Balance  string            `bson:"-" json:"balance"`
	Metadata *FungibleMetadata `bson:"metadata" json:"metadata"`
}

type FungibleMetadata struct {
	Name     string `bson:"name" json:"name"`
	Symbol   string `bson:"symbol" json:"symbol"`
	Decimals uint8  `bson:"decimals" json:"decimals"`
}
