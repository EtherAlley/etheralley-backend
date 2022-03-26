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

		switch fmt.Sprint(tokenId) {
		case common.STORE_PREMIUM:
			return &entities.NonFungibleMetadata{
				Name:        "Ether Alley Premium Membership",
				Description: "This token gives the holder access to premium features on EtherAlley.io",
				Image:       getImageUrl(settings.StoreImageURI(), "01"),
				Attributes: &[]map[string]interface{}{
					{"atr1": "val1", "atr2": "val2"},
				},
			}, nil
		case common.STORE_BETA_TESTER:
			return &entities.NonFungibleMetadata{
				Name:        "Ether Alley Beta Tester",
				Description: "The holder of this token participated in the Ether Alley beta. This token is non transferable",
				Image:       getImageUrl(settings.StoreImageURI(), "01"),
				Attributes: &[]map[string]interface{}{
					{"atr1": "val1", "atr2": "val2"},
				},
			}, nil
		}

		return nil, errors.New("unspported token id")
	}
}

func getImageUrl(baseUrl string, tokenId string) string {
	return fmt.Sprintf("%v/store/%v.png", baseUrl, tokenId)
}
