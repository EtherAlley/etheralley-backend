package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/etheralley/etheralley-backend/core/entities"
)

const ListingsNamespace = "listings"

func (g *gateway) GetStoreListings(ctx context.Context, tokenIds *[]string) (*[]entities.Listing, error) {

	listingsString, err := g.client.Get(ctx, getFullKey(ListingsNamespace, strings.Join(*tokenIds, "_"))).Result()

	if err != nil {
		return nil, fmt.Errorf("get listings %w", err)
	}

	listingsJson := &[]listingJson{}
	err = json.Unmarshal([]byte(listingsString), listingsJson)

	if err != nil {
		return nil, fmt.Errorf("get listings decoded %w", err)
	}

	listings := fromListingsJson(listingsJson)

	return listings, nil
}

func (g *gateway) SaveStoreListings(ctx context.Context, listings *[]entities.Listing) error {
	listingsJson := toListingsJson(listings)
	bytes, err := json.Marshal(listingsJson)

	if err != nil {
		return fmt.Errorf("save listings encode %w", err)
	}

	tokenIds := []string{}
	for _, listing := range *listings {
		tokenIds = append(tokenIds, listing.TokenId)
	}

	_, err = g.client.Set(ctx, getFullKey(ListingsNamespace, strings.Join(tokenIds, "_")), bytes, time.Hour).Result()

	if err != nil {
		return fmt.Errorf("save listings %w", err)
	}

	return nil
}
