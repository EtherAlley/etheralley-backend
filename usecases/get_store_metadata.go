package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

// Get the store metadata for the url that resolves from the contractURI call on the store contract
// See https://docs.opensea.io/docs/contract-level-metadata
type IGetStoreMetadataUseCase func(ctx context.Context) (metadata *entities.StoreMetadata)

func NewGetStoreMetadata(
	logger common.ILogger,
	settings common.ISettings,
) IGetStoreMetadataUseCase {
	return func(ctx context.Context) *entities.StoreMetadata {
		return &entities.StoreMetadata{
			Name:                 "EtherAlley Store",
			Description:          "Purchasable assets that unlock powerful features on the EtherAlley platform",
			Image:                "https://etheralley.io/store/contract.png",
			ExternalLink:         "https://etheralley.io",
			SellerFeeBasisPoints: 500, // 5%
			FeeRecipient:         "0x4c5d1C6eA7c485F491f9Dd79105949867b01fFa5",
		}
	}
}
