package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const viewLimitKey = "address_ipaddress"
const AddressViewsNamespace = "address_views"
const AddressViewsOut = "address_views_out"

func (gw *gateway) RecordAddressView(ctx context.Context, address string, ipAddress string) error {
	now := time.Now()
	bucket := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 9999, now.Location()) // get current hour bucket
	viewLimitTTL := bucket.Sub(now)

	viewLimitKey := getFullKey(viewLimitKey, address, ipAddress)
	addressViewsKey := getFullKey(AddressViewsNamespace, fmt.Sprint(bucket.Unix()))

	// if the viewer address combo exists, we return and do not record a profile view
	if err := gw.client.Get(ctx, viewLimitKey).Err(); err != redis.Nil {
		return err
	}

	if err := gw.client.Set(ctx, viewLimitKey, nil, viewLimitTTL).Err(); err != nil {
		return err
	}

	gw.client.ZIncrBy(ctx, addressViewsKey, 1, address).Err()

	return nil
}

func (gw *gateway) GetTopAddresses(ctx context.Context) (*[]string, error) {
	now := time.Now()
	keys := []string{}
	for i := 0; i < 24; i++ { // get buckets for the last 24 hours
		bucket := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()-i, 59, 59, 9999, now.Location())
		key := getFullKey(AddressViewsNamespace, fmt.Sprint(bucket.Unix()))
		keys = append(keys, key)
	}

	err := gw.client.ZUnionStore(ctx, AddressViewsOut, &redis.ZStore{
		Keys: keys,
	}).Err() // merge last 24 hour buckets into store

	if err != nil {
		return nil, err
	}

	results, err := gw.client.ZRevRangeWithScores(ctx, AddressViewsOut, 0, 9).Result() // take top 10 scores sorted in desc order

	if err != nil {
		return nil, err
	}

	topAddresss := []string{}
	for _, result := range results {
		topAddresss = append(topAddresss, result.Member.(string))
	}

	return &topAddresss, nil
}

// func (gw *gateway) RecordAddressView(ctx context.Context, address string, ipAddress string) error {
// 	now := time.Now()
// 	bucket := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 9999, now.Location())
// 	viewLimitTTL := bucket.Sub(now)
// 	addressViewsTTL := bucket.Sub(now) + time.Hour*23

// 	viewLimitKey := getFullKey(viewLimitKey, address, ipAddress)
// 	addressViewsKey := getFullKey(AddressViewsNamespace, address, fmt.Sprint(bucket.Unix()))

// 	// if the viewer address combo exists, we return and do not record a profile view
// 	if err := gw.client.Get(ctx, viewLimitKey).Err(); err != redis.Nil {
// 		return err
// 	}

// 	if err := gw.client.Set(ctx, viewLimitKey, nil, viewLimitTTL).Err(); err != nil {
// 		return err
// 	}

// 	if err := gw.client.Get(ctx, addressViewsKey).Err(); err == redis.Nil {
// 		gw.client.HSet(ctx, addressViewsKey, "count", 1).Err()
// 		gw.client.HSet(ctx, addressViewsKey, "address", address).Err()
// 		gw.client.Expire(ctx, addressViewsKey, addressViewsTTL).Err()
// 	} else {
// 		gw.client.HIncrBy(ctx, addressViewsKey, "count", 1)
// 	}

// 	return nil
// }

// func (gw *gateway) GetTopAddresses(ctx context.Context) (*[]string, error) {
// 	cursor := uint64(0)
// 	for {
// 		var keys []string
// 		var err error

// 		keys, cursor, err = gw.client.Scan(ctx, cursor, fmt.Sprintf("%v*", AddressViewsNamespace), 10).Result()

// 		fmt.Printf("%+v\n", keys)
// 		fmt.Printf("%+v\n", cursor)
// 		fmt.Printf("%+v\n", err)

// 		if cursor == 0 {
// 			break
// 		}
// 	}
// 	return nil, nil
// }

// func (gw *gateway) RecordAddressView2(ctx context.Context, address string, timestamp string, viewer string) error {
// 	viewLimitKey := getFullKey(viewLimitKey, address, viewer)
// 	addressViewsKey := getFullKey(AddressViewsNamespace, address)

// 	viewLimitTTL := 60
// 	addressViewsTTL := 600

// 	// result := gw.client.Incr(ctx, addressViewsKey)
// 	// fmt.Println(result)

// 	err := myIncrBy.Run(ctx, gw.client, []string{viewLimitKey, addressViewsKey}, viewLimitTTL, addressViewsTTL).Err()

// 	return err
// }

// var myIncrBy = redis.NewScript(`
// 	local viewLimitKey = KEYS[1]
// 	local addresViewsKey = KEYS[2]

// 	local viewLimitTTL = ARGV[1]
// 	local addressViewsTTL = ARGV[2]

// 	local hasViewed = redis.call("GET", viewLimitKey)

// 	if hasViewed then
// 		return
// 	end

// 	redis.call("SET", viewLimitKey, "", viewLimitTTL)

// 	local value = redis.call("GET", addresViewsKey)

// 	if not value then
// 		redis.call("SET", addresViewsKey, 1, addressViewsTTL)
// 		return
// 	end

// 	redis.call("INCR", addresViewsKey, "KEEPTTL")
// `)

// var incrBy = redis.NewScript(`
// 	local key = KEYS[1]
// 	local change = ARGV[1]
// 	local value = redis.call("GET", key)
// 	if not value then
// 		value = 0
// 	end
// 	value = value + change
// 	redis.call("SET", key, value)
// 	return value
// `)

// var sum = redis.NewScript(`
// 	local key = KEYS[1]
// 	local sum = redis.call("GET", key)
// 	if not sum then
// 		sum = 0
// 	end
// 	local num_arg = #ARGV
// 	for i = 1, num_arg do
// 		sum = sum + ARGV[i]
// 	end
// 	redis.call("SET", key, sum)
// 	return sum
// `)
