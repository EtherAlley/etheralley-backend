package entities

type NonFungibleToken struct {
	Contract *Contract `validate:"required"`
	TokenId  string    `validate:"required,numeric"`
	Balance  string
	Metadata *NonFungibleMetadata
}

type NonFungibleMetadata struct {
	Name        string
	Description string
	Image       string
	Attributes  *[]map[string]interface{}
}
