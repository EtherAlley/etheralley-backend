package entities

type NFT struct {
	TokenId         string      `bson:"token_id" json:"token_id"`
	Blockchain      string      `bson:"blockchain" json:"blockchain"`
	ContractAddress string      `bson:"address" json:"contract_address"`
	SchemaName      string      `bson:"schema_name" json:"schema_name"`
	Owned           bool        `bson:"-" json:"owned"`
	Metadata        NFTMetadata `bson:"metadata" json:"metadata"`
}

type NFTMetadata struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Image       string                   `json:"image"`
	Attributes  []map[string]interface{} `bson:"attributes" json:"attributes"`
	Properties  map[string]interface{}   `bson:"properties" json:"properties"`
}
