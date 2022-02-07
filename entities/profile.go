package entities

type Profile struct {
	Address           string              `bson:"_id" json:"-"`
	ENSName           string              `bson:"-" json:"ens_name"`
	NonFungibleTokens *[]NonFungibleToken `bson:"non_fungible_tokens" json:"non_fungible_tokens" validate:"required,dive"`
	FungibleTokens    *[]FungibleToken    `bson:"fungible_tokens" json:"fungible_tokens" validate:"required,dive"`
	Statistics        *[]Statistic        `bson:"statistics" json:"statistics" validate:"required,dive"`
}
