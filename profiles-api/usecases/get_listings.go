package usecases

import (
	"context"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
	"github.com/etheralley/etheralley-backend/profiles-api/gateways"
	"github.com/etheralley/etheralley-backend/profiles-api/settings"
)

func NewGetListings(
	logger common.ILogger,
	settings settings.ISettings,
	blockchainGateway gateways.IBlockchainGateway,
	cacheGateway gateways.ICacheGateway,
	getListingMetadata IGetListingMetadataUseCase,
) IGetListingsUseCase {
	return &getListingUseCase{
		logger,
		settings,
		blockchainGateway,
		cacheGateway,
		getListingMetadata,
	}
}

type getListingUseCase struct {
	logger             common.ILogger
	settings           settings.ISettings
	blockchainGateway  gateways.IBlockchainGateway
	cacheGateway       gateways.ICacheGateway
	getListingMetadata IGetListingMetadataUseCase
}

type IGetListingsUseCase interface {
	// Get the EtherAlley store listings for the provided array of token ids
	Do(ctx context.Context, input *GetListingsInput) (listings *[]entities.Listing, err error)
}

type GetListingsInput struct {
	TokenIds *[]string `json:"token_ids" validate:"required,dive,numeric"`
}

func (uc *getListingUseCase) Do(ctx context.Context, input *GetListingsInput) (*[]entities.Listing, error) {
	if err := common.ValidateStruct(input); err != nil {
		return nil, err
	}

	ids := *input.TokenIds

	cachedListings, err := uc.cacheGateway.GetStoreListings(ctx, &ids)

	if err == nil {
		uc.logger.Debug(ctx).Msgf("cache hit for store listings %+v", ids)
		return cachedListings, nil
	}

	uc.logger.Debug(ctx).Msgf("cache miss for store listings %+v", ids)

	listingInfo, err := uc.blockchainGateway.GetStoreListingInfo(ctx, &ids)

	if err != nil {
		uc.logger.Info(ctx).Err(err).Msgf("err getting store listings %+v", ids)
		return nil, err
	}

	listings := make([]entities.Listing, len(ids))
	for i := 0; i < len(ids); i++ {
		metadata, err := uc.getListingMetadata.Do(ctx, &GetListingMetadataInput{
			TokenId: ids[i],
		})

		if err != nil {
			return nil, err
		}

		listings[i] = entities.Listing{
			Contract: &entities.Contract{
				Blockchain: uc.settings.StoreBlockchain(),
				Address:    uc.settings.StoreAddress(),
				Interface:  common.ERC1155,
			},
			TokenId:  ids[i],
			Info:     &(*listingInfo)[i],
			Metadata: metadata,
		}
	}

	uc.cacheGateway.SaveStoreListings(ctx, &listings)

	return &listings, err
}
