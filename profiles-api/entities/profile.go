package entities

import "time"

type Profile struct {
	Address           string
	Banned            bool
	LastModified      *time.Time // LastModified can be nil if this is a defaul profile
	ENSName           string
	StoreAssets       *StoreAssets
	DisplayConfig     *DisplayConfig      // DisplayConfig can be nil if this is a default profile
	ProfilePicture    *NonFungibleToken   // ProfilePicture can be nil if this is a default profile with no NFTs or they have simply removed their profile picture
	NonFungibleTokens *[]NonFungibleToken // NonFungibleTokens can ve nil if this is a light hydrated profile
	FungibleTokens    *[]FungibleToken    // FungibleTokens can ve nil if this is a light hydrated profile
	Statistics        *[]Statistic        // Statistics can ve nil if this is a light hydrated profile
	Interactions      *[]Interaction      // Interactions can ve nil if this is a light hydrated profile
	Currencies        *[]Currency         // Currencies can ve nil if this is a light hydrated profile
}
