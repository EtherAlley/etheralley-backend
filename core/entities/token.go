package entities

type FungibleToken struct {
	Contract *Contract
	Balance  *string // balance can be nil
	Metadata *FungibleMetadata
}

type FungibleMetadata struct {
	Name     *string // name can be nil
	Symbol   *string // symbol can be nil
	Decimals *uint8  // decimals can be nil
}
