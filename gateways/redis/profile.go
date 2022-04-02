package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/go-redis/redis/v8"
)

const ProfileNamespace = "profile"

func (g *gateway) GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error) {
	profileString, err := g.client.Get(ctx, getFullKey(ProfileNamespace, address)).Result()

	if err == redis.Nil {
		return nil, common.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	profJson := &profileJson{}
	err = json.Unmarshal([]byte(profileString), profJson)

	if err != nil {
		return nil, err
	}

	profile := fromProfileJson(profJson)

	return profile, nil
}

func (g *gateway) SaveProfile(ctx context.Context, profile *entities.Profile) error {
	profJson := toProfileJson(profile)

	bytes, err := json.Marshal(profJson)

	if err != nil {
		return err
	}

	_, err = g.client.Set(ctx, getFullKey(ProfileNamespace, profile.Address), bytes, time.Hour*24).Result()

	return err
}
