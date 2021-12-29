package mongo

import (
	"context"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (g *Gateway) GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error) {
	profile := &entities.Profile{}
	err := g.profiles.FindOne(ctx, bson.D{{"_id", address}}).Decode(profile)

	if err == mongo.ErrNoDocuments {
		return profile, common.ErrNil
	}

	return profile, err
}

func (g *Gateway) SaveProfile(ctx context.Context, profile *entities.Profile) error {
	//TODO: Handle result
	_, err := g.profiles.UpdateOne(ctx, bson.D{{"_id", profile.Address}}, bson.D{{"$set", profile}}, options.Update().SetUpsert(true))

	return err
}
