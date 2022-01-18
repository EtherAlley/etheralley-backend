package opensea

import (
	"encoding/json"
	"fmt"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

type Gateway struct {
	logger   *common.Logger
	settings *common.Settings
	http     *common.HttpClient
}

func NewGateway(logger *common.Logger, settings *common.Settings, http *common.HttpClient) *Gateway {
	return &Gateway{
		logger,
		settings,
		http,
	}
}

type GetAssetsByOwnerRespBody struct {
	Assets []struct {
		TokenId       string                   `json:"token_id"`
		Image         string                   `json:"image_url"`
		Name          string                   `json:"name"`
		Description   string                   `json:"description"`
		Attributes    []map[string]interface{} `json:"traits"`
		AssetContract struct {
			ContractAddress string `json:"address"`
			SchemaName      string `json:"schema_name"`
		} `json:"asset_contract"`
	} `json:"assets"`
}

func (gw *Gateway) GetNonFungibleTokens(address string) *[]entities.NonFungibleToken {
	nfts := &[]entities.NonFungibleToken{}

	url := fmt.Sprintf("%v/assets?owner=%v&offset=0&limit=12", gw.settings.OpenSeaURI, address)
	resp, err := gw.http.Do("GET", url, &common.HttpOptions{
		Headers: []common.Header{
			{Key: "user-agent", Value: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"},
		},
	})

	if err != nil {
		gw.logger.Errf(err, "opensea assets get err: ")
		return nfts
	}

	defer resp.Body.Close()
	body := &GetAssetsByOwnerRespBody{}
	err = json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		gw.logger.Errf(err, "opensea assets decode err: ")
		return nfts
	}

	for _, asset := range body.Assets {
		attributes := asset.Attributes
		*nfts = append(*nfts, entities.NonFungibleToken{
			TokenId: asset.TokenId,
			Contract: &entities.Contract{
				Blockchain: common.ETHEREUM, // it appears that this API is only for layer 1
				Address:    asset.AssetContract.ContractAddress,
				Interface:  asset.AssetContract.SchemaName,
			},
			Balance: "1", //TODO: it doesnt appear opensea indicates the balance owned by the address, even if its semi-fungible?
			Metadata: &entities.NonFungibleMetadata{
				Name:        asset.Name,
				Description: asset.Description,
				Image:       asset.Image,
				Attributes:  &attributes,
			},
		})
	}

	return nfts
}
