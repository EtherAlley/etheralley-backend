package entities

type Profile struct {
	Address string `bson:"_id" json:"-"`
	NFTs    *[]NFT `bson:"nfts" json:"nfts" validate:"required,dive"`
}
