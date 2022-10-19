package entities

// See https://docs.opensea.io/docs/contract-level-metadata
type StoreMetadata struct {
	Name                 string
	Description          string
	Image                string
	ExternalLink         string
	SellerFeeBasisPoints uint
	FeeRecipient         string
}
