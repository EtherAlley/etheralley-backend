package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/etheralley/etheralley-backend/core/entities"
	"github.com/go-redis/redis/v8"
)

const ViewLimitNamespace = "top_profiles_view_limit"
const HourlyAddressViewsNamespace = "top_profiles_hourly_address_views"
const MergedAddressViews = "top_profiles_merged_address_views"
const TopProfilesResults = "top_profiles_results"

// hourly buckets are maintaned that collect views for a given address.
// if an address/ipaddress combo has already been counted in the hour, we do not increment the count for the address
func (gw *gateway) RecordAddressView(ctx context.Context, address string, ipAddress string) error {
	now := time.Now()
	bucket := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 9999, now.Location()) // get current hour bucket
	viewLimitTTL := bucket.Sub(now)

	viewLimitKey := getFullKey(ViewLimitNamespace, address, ipAddress)
	addressViewsKey := getFullKey(HourlyAddressViewsNamespace, fmt.Sprint(bucket.Unix()))

	// if the viewer address combo exists, we return and do not record a profile view
	if err := gw.client.Get(ctx, viewLimitKey).Err(); err != redis.Nil {
		return fmt.Errorf("view already counted %w", err)
	}

	// register the viewer/address combo as already having viewed
	if err := gw.client.Set(ctx, viewLimitKey, nil, viewLimitTTL).Err(); err != nil {
		return fmt.Errorf("setting view count %w", err)
	}

	gw.client.ZIncrBy(ctx, addressViewsKey, 1, address).Err() // increment address' view count within the zset

	// we unfortunately need to set the ttl on the bucket at some point...
	// the current bucket will live until its out of range of what GetTopAddresses will include in its selection
	if ttl, err := gw.client.TTL(ctx, addressViewsKey).Result(); err == nil && ttl == -1 {
		gw.client.Expire(ctx, addressViewsKey, viewLimitTTL+(time.Hour*24)).Err()
	}

	return nil
}

func (gw *gateway) GetTopViewedAddresses(ctx context.Context) (*[]string, error) {
	// get buckets for the last 24 hours
	now := time.Now()
	keys := []string{}
	for i := 0; i < 24; i++ {
		bucket := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()-i, 59, 59, 9999, now.Location())
		key := getFullKey(HourlyAddressViewsNamespace, fmt.Sprint(bucket.Unix()))
		keys = append(keys, key)
	}

	// merge last 24 hour buckets into store
	if err := gw.client.ZUnionStore(ctx, MergedAddressViews, &redis.ZStore{
		Keys: keys,
	}).Err(); err != nil {
		return nil, fmt.Errorf("merging buckets %w", err)
	}

	// take top 10 scores sorted in desc order
	// TODO: can parameterize this in the future
	results, err := gw.client.ZRevRangeWithScores(ctx, MergedAddressViews, 0, 9).Result()

	if err != nil {
		return nil, fmt.Errorf("sorting scores %w", err)
	}

	topAddresses := []string{}
	for _, result := range results {
		topAddresses = append(topAddresses, result.Member.(string))
	}

	return &topAddresses, nil
}

func (gw *gateway) GetTopViewedProfiles(ctx context.Context) (*[]entities.Profile, error) {
	profilesString, err := gw.client.Get(ctx, TopProfilesResults).Result()

	if err != nil {
		return nil, fmt.Errorf("get top viewed profiles %w", err)
	}

	results := &[]profileJson{}
	err = json.Unmarshal([]byte(profilesString), results)

	if err != nil {
		return nil, fmt.Errorf("decode top viewed profiles %w", err)
	}

	profiles := []entities.Profile{}

	for _, result := range *results {
		profiles = append(profiles, *fromProfileJson(&result))
	}

	return &profiles, nil
}

func (gw *gateway) SaveTopViewedProfiles(ctx context.Context, profiles *[]entities.Profile) error {
	profilesJson := []profileJson{}

	for _, profile := range *profiles {
		profilesJson = append(profilesJson, *toProfileJson(&profile))
	}

	bytes, err := json.Marshal(&profilesJson)

	if err != nil {
		return fmt.Errorf("encode top viewed profiles %w", err)
	}

	err = gw.client.Set(ctx, TopProfilesResults, bytes, time.Hour).Err()

	if err != nil {
		return fmt.Errorf("set top viewed profiles %w", err)
	}

	return nil
}
