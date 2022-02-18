package entities

type Profile struct {
	Address           string
	ENSName           string
	NonFungibleTokens *[]NonFungibleToken
	FungibleTokens    *[]FungibleToken
	Statistics        *[]Statistic
	Interactions      *[]Interaction
}
