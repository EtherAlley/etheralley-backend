package redis

import (
	"context"
	"fmt"
	"time"
)

const ENSNamespace = "ens"

func (g *gateway) GetENSAddressFromName(ctx context.Context, ensName string) (string, error) {
	address, err := g.client.Get(ctx, getFullKey(ENSNamespace, ensName)).Result()

	if err != nil {
		return "", fmt.Errorf("get ens address %w", err)
	}

	return address, nil
}

func (g *gateway) SaveENSAddress(ctx context.Context, ensName string, address string) error {
	_, err := g.client.Set(ctx, getFullKey(ENSNamespace, ensName), address, time.Hour*24).Result()

	if err != nil {
		return fmt.Errorf("save ens address %w", err)
	}

	return nil
}

func (g *gateway) GetENSNameFromAddress(ctx context.Context, address string) (string, error) {
	name, err := g.client.Get(ctx, getFullKey(ENSNamespace, address)).Result()

	if err != nil {
		return "", fmt.Errorf("get ens name %w", err)
	}

	return name, nil
}

func (g *gateway) SaveENSName(ctx context.Context, address string, name string) error {
	_, err := g.client.Set(ctx, getFullKey(ENSNamespace, address), name, time.Hour*24).Result()

	if err != nil {
		return fmt.Errorf("save ens name %w", err)
	}

	return nil
}
