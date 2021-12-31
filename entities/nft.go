package entities

type NFT struct {
	TokenId         string              `bson:"token_id" json:"token_id"`
	Blockchain      string              `bson:"blockchain" json:"blockchain"`
	ContractAddress string              `bson:"address" json:"contract_address"`
	SchemaName      string              `bson:"schema_name" json:"schema_name"`
	Owned           bool                `bson:"-" json:"owned"`
	Name            string              `bson:"name" json:"name"`
	Description     string              `bson:"description" json:"description"`
	ImageURL        string              `bson:"image_url" json:"image_url"`
	Attributes      []map[string]string `bson:"attributes" json:"attributes"`
}
