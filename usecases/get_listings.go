package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type GetListingsInput struct {
	TokenIds *[]string `json:"token_ids" validate:"required,dive,numeric"`
}

type IGetListingsUseCase func(ctx context.Context, input *GetListingsInput) (listings *[]entities.Listing, err error)

func NewGetListings(
	logger common.ILogger,
	settings common.ISettings,
	blockchainGateway gateways.IBlockchainGateway,
	getListingMetadata IGetListingMetadataUseCase,
) IGetListingsUseCase {
	return func(ctx context.Context, input *GetListingsInput) (*[]entities.Listing, error) {
		if err := common.ValidateStruct(input); err != nil {
			return nil, err
		}

		ids := *input.TokenIds

		listingInfo, err := blockchainGateway.GetStoreListingInfo(ctx, &ids)

		if err != nil {
			return nil, err
		}

		listings := make([]entities.Listing, len(ids))
		for i := 0; i < len(ids); i++ {
			metadata, err := getListingMetadata(ctx, &GetListingMetadataInput{
				TokenId: ids[i],
			})

			if err != nil {
				return nil, err
			}

			listings[i] = entities.Listing{
				Contract: &entities.Contract{
					Blockchain: settings.StoreBlockchain(),
					Address:    settings.StoreAddress(),
					Interface:  common.ERC1155,
				},
				TokenId:  ids[i],
				Info:     &(*listingInfo)[i],
				Metadata: metadata,
			}
		}

		return &listings, err
	}
}
