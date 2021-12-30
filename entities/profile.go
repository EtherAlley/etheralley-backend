package entities

type Profile struct {
	Address     string       `bson:"_id" json:"-"`
	NFTElements []NFTElement `bson:"nft_elements" json:"nft_elements"`
}

type NFTElement struct {
	Order      uint   `bson:"order" json:"order"`
	Address    string `bson:"address" json:"address"`
	SchemaName string `bson:"schema_name" json:"schema_name"`
	TokenId    string `bson:"token_id" json:"token_id"`
}
