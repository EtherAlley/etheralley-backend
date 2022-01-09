package entities

type NFT struct {
	Owned    bool         `bson:"-" json:"owned"`
	Location *NFTLocation `bson:"location" json:"location" validate:"required"`
	Metadata *NFTMetadata `bson:"metadata" json:"metadata"`
}

type NFTLocation struct {
	TokenId         string `bson:"token_id" json:"token_id" validate:"required,numeric"`
	Blockchain      string `bson:"blockchain" json:"blockchain" validate:"required,oneof=ethereum polygon arbitrum optimism"`
	ContractAddress string `bson:"contract_address" json:"contract_address" validate:"required,eth_addr"`
	SchemaName      string `bson:"schema_name" json:"schema_name" validate:"required,oneof=ERC721 ERC1155"`
}

type NFTMetadata struct {
	Name        string                    `bson:"name" json:"name"`
	Description string                    `bson:"description" json:"description"`
	Image       string                    `bson:"image" json:"image"`
	Attributes  *[]map[string]interface{} `bson:"attributes" json:"attributes"`
	Properties  *map[string]interface{}   `bson:"properties" json:"properties"`
}
