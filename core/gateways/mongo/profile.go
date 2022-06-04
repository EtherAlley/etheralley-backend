package mongo

import (
	"context"
	"fmt"

	"github.com/etheralley/etheralley-apis/common"
	"github.com/etheralley/etheralley-apis/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (g *gateway) GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error) {
	profileBson := &profileBson{}

	err := g.profiles.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: address}}).Decode(profileBson)

	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("profile not found %w", common.ErrNotFound)
	}

	if err != nil {
		return nil, fmt.Errorf("get profile %w", err)
	}

	profile := fromProfileBson(profileBson)

	return profile, nil
}

func (g *gateway) SaveProfile(ctx context.Context, profile *entities.Profile) error {
	profileBson := toProfileBson(profile)

	_, err := g.profiles.UpdateOne(ctx, bson.D{primitive.E{Key: "_id", Value: profile.Address}}, bson.D{primitive.E{Key: "$set", Value: profileBson}}, options.Update().SetUpsert(true))

	if err != nil {
		return fmt.Errorf("save profile %w", err)
	}

	return nil
}
