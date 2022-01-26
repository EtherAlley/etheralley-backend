package entities

type NonFungibleToken struct {
	Contract *Contract            `bson:"contract" json:"contract" validate:"required"`
	TokenId  string               `bson:"token_id" json:"token_id" validate:"required,numeric"`
	Balance  string               `json:"balance"`
	Metadata *NonFungibleMetadata `json:"metadata"`
}

type NonFungibleMetadata struct {
	Name        string                    `bson:"name" json:"name"`
	Description string                    `bson:"description" json:"description"`
	Image       string                    `bson:"image" json:"image"`
	Attributes  *[]map[string]interface{} `bson:"attributes" json:"attributes"`
}
