package entities

type Profile struct {
	Address           string
	ENSName           string
	NonFungibleTokens *[]NonFungibleToken `validate:"required,dive"`
	FungibleTokens    *[]FungibleToken    `validate:"required,dive"`
	Statistics        *[]Statistic        `validate:"required,dive"`
	Interactions      *[]Interaction      `validate:"required,dive"`
}
