package redis

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/etheralley/etheralley-core-api/entities"
)

const ListingsNamespace = "listings"

func (g *gateway) GetStoreListings(ctx context.Context, tokenIds *[]string) (*[]entities.Listing, error) {

	listingsString, err := g.client.Get(ctx, getFullKey(ListingsNamespace, strings.Join(*tokenIds, "_"))).Result()

	if err != nil {
		return nil, err
	}

	listingsJson := &[]listingJson{}
	err = json.Unmarshal([]byte(listingsString), listingsJson)

	if err != nil {
		return nil, err
	}

	listings := fromListingsJson(listingsJson)

	return listings, nil
}

func (g *gateway) SaveStoreListings(ctx context.Context, listings *[]entities.Listing) error {
	listingsJson := toListingsJson(listings)
	bytes, err := json.Marshal(listingsJson)

	if err != nil {
		return err
	}

	tokenIds := []string{}
	for _, listing := range *listings {
		tokenIds = append(tokenIds, listing.TokenId)
	}

	_, err = g.client.Set(ctx, getFullKey(ListingsNamespace, strings.Join(tokenIds, "_")), bytes, time.Hour).Result()

	return err
}
