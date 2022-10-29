package entities

import "time"

type Profile struct {
	Address           string
	Banned            bool
	LastModified      *time.Time // LastModified can be nil if this is a defaul profile
	ENSName           string
	StoreAssets       *StoreAssets
	DisplayConfig     *DisplayConfig    // DisplayConfig can be nil if this is a default profile
	ProfilePicture    *NonFungibleToken // ProfilePicture can be nil if this is a default profile with no NFTs or they have simply removed their profile picture
	NonFungibleTokens *[]NonFungibleToken
	FungibleTokens    *[]FungibleToken
	Statistics        *[]Statistic
	Interactions      *[]Interaction
	Currencies        *[]Currency
}
