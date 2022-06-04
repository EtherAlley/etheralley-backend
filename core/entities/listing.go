package entities

type Listing struct {
	Contract *Contract
	TokenId  string
	Info     *ListingInfo
	Metadata *NonFungibleMetadata
}

type ListingInfo struct {
	Purchasable  bool
	Transferable bool
	Price        string
	BalanceLimit string
	SupplyLimit  string
}
