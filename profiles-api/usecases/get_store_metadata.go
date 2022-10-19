package usecases

import (
	"context"

	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

func NewGetStoreMetadata() IGetStoreMetadataUseCase {
	return &getStoreMetadataUseCase{}
}

type getStoreMetadataUseCase struct {
}

type IGetStoreMetadataUseCase interface {
	// Get the store metadata for the url that resolves from the contractURI call on the store contract
	Do(ctx context.Context) (metadata *entities.StoreMetadata)
}

// See https://docs.opensea.io/docs/contract-level-metadata
func (uc *getStoreMetadataUseCase) Do(ctx context.Context) *entities.StoreMetadata {
	return &entities.StoreMetadata{
		Name:                 "EtherAlley Store",
		Description:          "Purchasable assets that unlock powerful features on the EtherAlley.io platform",
		Image:                "https://etheralley.io/store/contract.png",
		ExternalLink:         "https://etheralley.io",
		SellerFeeBasisPoints: 500, // 5%
		FeeRecipient:         "0x4c5d1C6eA7c485F491f9Dd79105949867b01fFa5",
	}
}
