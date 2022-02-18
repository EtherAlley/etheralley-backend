package entities

type FungibleToken struct {
	Contract *Contract
	Balance  string
	Metadata *FungibleMetadata
}

type FungibleMetadata struct {
	Name     string
	Symbol   string
	Decimals uint8
}
