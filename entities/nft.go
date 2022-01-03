package entities

type NFT struct {
	Owned    bool         `bson:"-" json:"owned"`
	Location *NFTLocation `bson:"location" json:"location"`
	Metadata *NFTMetadata `bson:"metadata" json:"metadata"`
}

type NFTLocation struct {
	TokenId         string `bson:"token_id" json:"token_id"`
	Blockchain      string `bson:"blockchain" json:"blockchain"`
	ContractAddress string `bson:"contract_address" json:"contract_address"`
	SchemaName      string `bson:"schema_name" json:"schema_name"`
}

type NFTMetadata struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Image       string                    `json:"image"`
	Attributes  *[]map[string]interface{} `bson:"attributes" json:"attributes"`
	Properties  *map[string]interface{}   `bson:"properties" json:"properties"`
}
