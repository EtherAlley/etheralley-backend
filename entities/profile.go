package entities

type Profile struct {
	Address           string              `bson:"_id" json:"-"`
	NonFungibleTokens *[]NonFungibleToken `bson:"non_fungible_tokens" json:"non_fungible_tokens" validate:"required,dive"`
	FungibleTokens    *[]FungibleToken    `bson:"fungible_tokens" json:"fungible_tokens" validate:"required,dive"`
}
