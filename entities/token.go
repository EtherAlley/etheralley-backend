package entities

type FungibleToken struct {
	Contract *Contract `validate:"required"`
	Balance  string
	Metadata *FungibleMetadata
}

type FungibleMetadata struct {
	Name     string
	Symbol   string
	Decimals uint8
}
