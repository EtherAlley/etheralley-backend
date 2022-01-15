package entities

type NonFungibleToken struct {
	Contract *Contract            `bson:"contract" json:"contract" validate:"required"`
	TokenId  string               `bson:"token_id" json:"token_id" validate:"required,numeric"`
	Balance  string               `json:"balance"`
	Metadata *NonFungibleMetadata `json:"metadata"`
}

type FungibleToken struct {
	Contract *Contract         `bson:"contract" json:"contract" validate:"required"`
	Balance  string            `json:"balance"`
	Metadata *FungibleMetadata `json:"metadata"`
}

type Contract struct {
	Blockchain string `bson:"blockchain" json:"blockchain" validate:"required,oneof=ethereum polygon arbitrum optimism"`
	Address    string `bson:"address" json:"address" validate:"required,eth_addr"`
	Interface  string `bson:"interface" json:"interface" validate:"required,oneof=ERC721 ERC1155 ERC20 ENS_REGISTRAR"`
}

type NonFungibleMetadata struct {
	Name        string                    `bson:"name" json:"name"`
	Description string                    `bson:"description" json:"description"`
	Image       string                    `bson:"image" json:"image"`
	Attributes  *[]map[string]interface{} `bson:"attributes" json:"attributes"`
}

type FungibleMetadata struct {
	Name     string `bson:"name" json:"name"`
	Symbol   string `bson:"symbol" json:"symbol"`
	Decimals uint8  `bson:"decimals" json:"decimals"`
}
