package entities

import "time"

type Profile struct {
	Address           string
	Banned            bool
	LastModified      *time.Time // LastModified can be nil if this is a defaul profile
	ENSName           string
	StoreAssets       *StoreAssets
	DisplayConfig     *DisplayConfig    // DisplayConfig can be nil if this is a default profile
	ProfilePicture    *NonFungibleToken // can be nil if this is a default profile
	NonFungibleTokens *[]NonFungibleToken
	FungibleTokens    *[]FungibleToken
	Statistics        *[]Statistic
	Interactions      *[]Interaction
	Currencies        *[]Currency
}
