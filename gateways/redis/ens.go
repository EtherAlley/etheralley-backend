package redis

import (
	"context"
	"time"
)

const ENSNamespace = "ens"

func (g *Gateway) GetENSAddressFromName(ctx context.Context, ensName string) (string, error) {
	return g.client.Get(ctx, GetFullKey(ENSNamespace, ensName)).Result()
}

func (g *Gateway) SaveENSAddress(ctx context.Context, ensName string, address string) error {
	_, err := g.client.Set(ctx, GetFullKey(ENSNamespace, ensName), address, time.Hour*24).Result()
	return err
}
