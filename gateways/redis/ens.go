package redis

import (
	"context"
	"time"
)

const ENSNamespace = "ens"

func (g *Gateway) GetENSAddressFromName(ctx context.Context, ensName string) (string, error) {
	return g.client.Get(ctx, getFullKey(ENSNamespace, ensName)).Result()
}

func (g *Gateway) SaveENSAddress(ctx context.Context, ensName string, address string) error {
	_, err := g.client.Set(ctx, getFullKey(ENSNamespace, ensName), address, time.Hour*24).Result()
	return err
}
