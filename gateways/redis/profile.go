package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/etheralley/etheralley-core-api/entities"
)

const ProfileNamespace = "profile"

func (g *Gateway) GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error) {
	profileString, err := g.client.Get(ctx, getFullKey(ProfileNamespace, address)).Result()

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

func (g *Gateway) SaveProfile(ctx context.Context, profile *entities.Profile) error {
	profJson := toProfileJson(profile)

	bytes, err := json.Marshal(profJson)

	if err != nil {
		return err
	}

	_, err = g.client.Set(ctx, getFullKey(ProfileNamespace, profile.Address), bytes, time.Hour*24).Result()

	return err
}

func fromProfileJson(profileJson *profileJson) *entities.Profile {
	profile := &entities.Profile{
		Address:           profileJson.Address,
		ENSName:           profileJson.ENSName,
		NonFungibleTokens: &[]entities.NonFungibleToken{},
		FungibleTokens:    &[]entities.FungibleToken{},
		Statistics:        &[]entities.Statistic{},
		Interactions:      &[]entities.Interaction{},
	}
	return profile
}

func toProfileJson(profile *entities.Profile) *profileJson {
	profileJson := &profileJson{
		Address:           profile.Address,
		ENSName:           profile.ENSName,
		NonFungibleTokens: &[]nonFungibleTokenJson{},
		FungibleTokens:    &[]fungibleTokenJson{},
		Statistics:        &[]statisticJson{},
		Interactions:      &[]interactionJson{},
	}
	return profileJson
}
