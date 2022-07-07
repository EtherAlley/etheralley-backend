package offchain

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/etheralley/etheralley-backend/core/entities"
)

type kittieRespBody struct {
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	ImageURL string `json:"image_url_cdn"`
	Matron   struct {
		EnhancedCattributes []struct {
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"enhanced_cattributes"`
	} `json:"matron"`
}

func (gw *gateway) GetKittieMetadata(ctx context.Context, tokenId string) (*entities.NonFungibleMetadata, error) {
	uri := fmt.Sprintf("%v/%v", gw.settings.CryptoKittiesMetadataURI(), tokenId)

	resp, err := gw.httpClient.Do(ctx, "GET", uri, nil, nil)

	if err != nil {
		return nil, fmt.Errorf("crypto kittie get metadata %w", err)
	}

	metadata := &kittieRespBody{}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&metadata)

	if err != nil {
		return nil, fmt.Errorf("crypto kittie decode metadata %w", err)
	}

	attributes := []map[string]interface{}{}
	for _, cattribute := range metadata.Matron.EnhancedCattributes {
		attributes = append(attributes, map[string]interface{}{"trait_type": cattribute.Type, "value": cattribute.Description})
	}

	return &entities.NonFungibleMetadata{
		Name:        metadata.Name,
		Description: metadata.Bio,
		Image:       metadata.ImageURL,
		Attributes:  &attributes,
	}, nil
}
