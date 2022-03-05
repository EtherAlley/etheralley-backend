package entities

type NonFungibleToken struct {
	Contract *Contract
	TokenId  string
	Balance  string
	Metadata *NonFungibleMetadata // Metadata can be nil
}

type NonFungibleMetadata struct {
	Name        string
	Description string
	Image       string
	Attributes  *[]map[string]interface{}
}
