package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/etheralley/etheralley-core-api/entities"
)

const ProfileNamespace = "profile_"

func (g *Gateway) GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error) {
	profile := &entities.Profile{}

	profileString, err := g.client.Get(ctx, GetFullKey(ProfileNamespace, address)).Result()

	if err != nil {
		return profile, err
	}

	err = json.Unmarshal([]byte(profileString), &profile)

	return profile, err
}

func (g *Gateway) SaveProfile(ctx context.Context, profile *entities.Profile) error {
	profileBytes, err := json.Marshal(profile)

	if err != nil {
		return err
	}

	_, err = g.client.Set(ctx, GetFullKey(ProfileNamespace, profile.Address), string(profileBytes), time.Hour).Result()

	return err
}
