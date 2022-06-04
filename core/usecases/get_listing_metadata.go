package usecases

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/core/entities"
	"github.com/etheralley/etheralley-core-api/core/settings"
)

func NewGetListingMetadata(
	logger common.ILogger,
	settings settings.ISettings,
) IGetListingMetadataUseCase {
	return &getListingMetadata{
		logger,
		settings,
	}
}

type getListingMetadata struct {
	logger   common.ILogger
	settings settings.ISettings
}

type IGetListingMetadataUseCase interface {
	// Get the metadata for a provided token id for the EtherAlley store
	Do(ctx context.Context, input *GetListingMetadataInput) (metadata *entities.NonFungibleMetadata, err error)
}

type GetListingMetadataInput struct {
	TokenId string `validate:"required,numeric"`
}

func (uc *getListingMetadata) Do(ctx context.Context, input *GetListingMetadataInput) (*entities.NonFungibleMetadata, error) {
	if err := common.ValidateStruct(input); err != nil {
		return nil, err
	}

	// See https://eips.ethereum.org/EIPS/eip-1155: token ids are passed in hexadecimal form with no 0x prefix.
	tokenId, err := strconv.ParseInt(input.TokenId, 16, 64)

	if err != nil {
		return nil, err
	}

	tokenString := fmt.Sprint(tokenId)
	switch tokenString {
	case common.STORE_PREMIUM:
		return &entities.NonFungibleMetadata{
			Name:        "EtherAlley Premium",
			Description: "This is a semi-fungible token that gives the holder access to premium features on EtherAlley.io",
			Image:       uc.getImageUrl(tokenString),
			Attributes: getAttribute([][2]interface{}{
				{"Status", "Verified"},
				{"Badge Count", "50"},
				{"Fungibility", "Semi-Fungible"},
				{"Transferable", "true"},
				{"Max balance", "unlimited"},
			}),
		}, nil
	case common.STORE_BETA_TESTER:
		return &entities.NonFungibleMetadata{
			Name:        "EtherAlley Beta Tester",
			Description: "This is a semi-fungible & soulbound token that indicates the holder participated in the EtherAlley.io beta",
			Image:       uc.getImageUrl(tokenString),
			Attributes: getAttribute([][2]interface{}{
				{"Achievement", "Beta Tester"},
				{"Fungibility", "Semi-Fungible"},
				{"Transferable", "false"},
				{"Max balance", "1"},
			}),
		}, nil
	}

	return nil, errors.New("unspported token id")
}

func (uc *getListingMetadata) getImageUrl(tokenId string) string {
	return fmt.Sprintf("%v/store/%v.png", uc.settings.StoreImageURI(), tokenId)
}

func getAttribute(attrs [][2]interface{}) *[]map[string]interface{} {
	attributes := []map[string]interface{}{}

	for _, attribute := range attrs {
		attributes = append(attributes, map[string]interface{}{
			"trait_type": attribute[0], "value": attribute[1],
		})
	}

	return &attributes
}
