package opensea

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

type Gateway struct {
	logger *common.Logger
}

func NewGateway(logger *common.Logger) *Gateway {
	return &Gateway{
		logger,
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

const OpenSeaBaseUrl = "https://api.opensea.io/api/v1"

func (gw *Gateway) GetNFTs(address string) ([]entities.NFT, error) {
	nfts := []entities.NFT{}

	url := fmt.Sprintf("%v/assets?owner=%v&offset=0&limit=50", OpenSeaBaseUrl, address)

	gw.logger.Info(url)

	resp, err := http.Get(url)

	if err != nil {
		gw.logger.Errf(err, "open sea gateway get nfts http call: ")
		return nfts, nil
	}

	if resp.StatusCode != 200 {
		gw.logger.Errorf("open sea gateway get nfts http call: %v", resp.StatusCode)
		return nfts, nil
	}

	defer resp.Body.Close()
	body := &GetAssetsByOwnerRespBody{}
	err = json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		gw.logger.Errf(err, "open sea gateway get nfts decode response body: ")
		return nfts, nil
	}

	for _, asset := range body.Assets {
		nfts = append(nfts, entities.NFT{
			TokenId:         asset.TokenId,
			Blockchain:      "ethereum", // TODO: ???
			ContractAddress: asset.AssetContract.ContractAddress,
			SchemaName:      asset.AssetContract.SchemaName,
			Owned:           true,
			Metadata: entities.NFTMetadata{
				Name:        asset.Name,
				Description: asset.Description,
				Image:       asset.Image,
				Attributes:  asset.Attributes,
			},
		})
	}

	return nfts, nil
}
