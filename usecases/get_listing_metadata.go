package usecases

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

type GetListingMetadataInput struct {
	TokenId string `validate:"required,numeric"`
}

// Get the metadata for a provided token id for the EtherAlley store
type IGetListingMetadataUseCase func(ctx context.Context, input *GetListingMetadataInput) (metadata *entities.NonFungibleMetadata, err error)

func NewGetListingMetadata(
	logger common.ILogger,
	settings common.ISettings,
) IGetListingMetadataUseCase {
	return func(ctx context.Context, input *GetListingMetadataInput) (*entities.NonFungibleMetadata, error) {
		if err := common.ValidateStruct(input); err != nil {
			return nil, err
		}

		// See https://eips.ethereum.org/EIPS/eip-1155: token ids are passed in hexidecimal form
		tokenId, err := strconv.ParseInt(input.TokenId, 16, 64)

		if err != nil {
			return nil, err
		}

		tokenString := fmt.Sprint(tokenId)

		url := getImageUrl(settings.StoreImageURI(), tokenString)

		switch tokenString {
		case common.STORE_PREMIUM:
			return &entities.NonFungibleMetadata{
				Name:        "EtherAlley Premium",
				Description: "This token gives the holder access to premium features on EtherAlley.io",
				Image:       url,
				Attributes: getAttribute([][2]interface{}{
					{"Status", "Verified"},
					{"Fungibility", "Semi-Fungible"},
					{"Transferable", "True"},
				}),
			}, nil
		case common.STORE_BETA_TESTER:
			return &entities.NonFungibleMetadata{
				Name:        "EtherAlley Beta Tester",
				Description: "The holder of this token participated in the EtherAlley beta. This token is non transferable",
				Image:       url,
				Attributes: getAttribute([][2]interface{}{
					{"Achievement", "Beta Tester"},
					{"Fungibility", "Semi-Fungible"},
					{"Transferable", "False"},
					{"Maximum balance per address", "1"},
				}),
			}, nil
		}

		return nil, errors.New("unspported token id")
	}
}

func getImageUrl(baseUrl string, tokenId string) string {
	return fmt.Sprintf("%v/store/%v.png", baseUrl, tokenId)
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
