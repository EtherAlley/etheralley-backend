package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
	"github.com/go-redis/redis/v8"
)

const ProfileNamespace = "profiles"

func (g *gateway) GetProfileByAddress(ctx context.Context, key string, address string) (*entities.Profile, error) {
	profileString, err := g.client.Get(ctx, getFullKey(ProfileNamespace, key, address)).Result()

	if err == redis.Nil {
		return nil, fmt.Errorf("profile %v not found %w", key, common.ErrNotFound)
	}

	if err != nil {
		return nil, fmt.Errorf("get profile %v %w", key, err)
	}

	profJson := &profileJson{}
	err = json.Unmarshal([]byte(profileString), profJson)

	if err != nil {
		return nil, fmt.Errorf("decode profile %v %w", key, err)
	}

	profile := fromProfileJson(profJson)

	return profile, nil
}

func (g *gateway) SaveProfile(ctx context.Context, key string, profile *entities.Profile) error {
	profJson := toProfileJson(profile)

	bytes, err := json.Marshal(profJson)

	if err != nil {
		return fmt.Errorf("encode profile %v %w", key, err)
	}

	_, err = g.client.Set(ctx, getFullKey(ProfileNamespace, key, profile.Address), bytes, time.Hour).Result()

	if err != nil {
		return fmt.Errorf("save profile %v %w", key, err)
	}

	return nil
}

func (g *gateway) DeleteProfile(ctx context.Context, key string, address string) error {
	err := g.client.Del(ctx, getFullKey(ProfileNamespace, key, address)).Err()

	if err != nil {
		return fmt.Errorf("delete profile %v %w", key, err)
	}

	return nil
}

func (gw *gateway) GetProfiles(ctx context.Context, key string) (*[]entities.Profile, error) {
	profilesString, err := gw.client.Get(ctx, getFullKey(ProfileNamespace, key)).Result()

	if err != nil {
		return nil, fmt.Errorf("get profiles %v %w", key, err)
	}

	results := &[]profileJson{}
	err = json.Unmarshal([]byte(profilesString), results)

	if err != nil {
		return nil, fmt.Errorf("decode profiles %v %w", key, err)
	}

	profiles := []entities.Profile{}

	for _, result := range *results {
		profiles = append(profiles, *fromProfileJson(&result))
	}

	return &profiles, nil
}

func (gw *gateway) SaveProfiles(ctx context.Context, key string, profiles *[]entities.Profile) error {
	profilesJson := []profileJson{}

	for _, profile := range *profiles {
		profilesJson = append(profilesJson, *toProfileJson(&profile))
	}

	bytes, err := json.Marshal(&profilesJson)

	if err != nil {
		return fmt.Errorf("encode profiles %v %w", key, err)
	}

	err = gw.client.Set(ctx, getFullKey(ProfileNamespace, key), bytes, time.Hour).Err()

	if err != nil {
		return fmt.Errorf("set profiles %v %w", key, err)
	}

	return nil
}
