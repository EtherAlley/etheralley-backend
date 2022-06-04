package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/core/entities"
	"github.com/go-redis/redis/v8"
)

const ProfileNamespace = "profile"

func (g *gateway) GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error) {
	profileString, err := g.client.Get(ctx, getFullKey(ProfileNamespace, address)).Result()

	if err == redis.Nil {
		return nil, fmt.Errorf("profile not found %w", common.ErrNotFound)
	}

	if err != nil {
		return nil, fmt.Errorf("get profile %w", err)
	}

	profJson := &profileJson{}
	err = json.Unmarshal([]byte(profileString), profJson)

	if err != nil {
		return nil, fmt.Errorf("decode profile %w", err)
	}

	profile := fromProfileJson(profJson)

	return profile, nil
}

func (g *gateway) SaveProfile(ctx context.Context, profile *entities.Profile) error {
	profJson := toProfileJson(profile)

	bytes, err := json.Marshal(profJson)

	if err != nil {
		return fmt.Errorf("encode profile %w", err)
	}

	_, err = g.client.Set(ctx, getFullKey(ProfileNamespace, profile.Address), bytes, time.Hour).Result()

	if err != nil {
		return fmt.Errorf("save profile %w", err)
	}

	return nil
}

func (g *gateway) DeleteProfile(ctx context.Context, address string) error {
	err := g.client.Del(ctx, getFullKey(ProfileNamespace, address)).Err()

	if err != nil {
		return fmt.Errorf("delete profile %w", err)
	}

	return nil
}
