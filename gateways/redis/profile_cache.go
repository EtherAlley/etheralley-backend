package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/eflem00/go-example-app/entities"
)

type ProfileCache struct {
	cache *Cache
}

func NewProfileCache(cache *Cache) *ProfileCache {
	return &ProfileCache{
		cache,
	}
}

const ProfileNamespace = "profile_"

func (pc *ProfileCache) GetProfileByAddress(ctx context.Context, address string) (entities.Profile, error) {
	profile := entities.Profile{}

	profileString, err := pc.cache.Client.Get(ctx, GetFullKey(ProfileNamespace, address)).Result()

	if err != nil {
		return profile, err
	}

	err = json.Unmarshal([]byte(profileString), &profile)

	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (pc *ProfileCache) SaveProfile(ctx context.Context, profile entities.Profile) error {
	profileBytes, err := json.Marshal(profile)

	if err != nil {
		return err
	}

	_, err = pc.cache.Client.Set(ctx, GetFullKey(ProfileNamespace, profile.Address), string(profileBytes), time.Hour).Result()

	return err
}
