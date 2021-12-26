package mongo

import (
	"context"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: fill in the persistant storage piece

type ProfileRepository struct {
	profiles *mongo.Collection
	logger   *common.Logger
}

func NewProfileRepository(db *mongo.Database, logger *common.Logger) *ProfileRepository {
	profiles := db.Collection("profiles")
	return &ProfileRepository{
		profiles,
		logger,
	}
}

func (repo *ProfileRepository) GetProfileByAddress(ctx context.Context, address string) (entities.Profile, error) {
	var profile entities.Profile
	err := repo.profiles.FindOne(ctx, bson.D{{"_id", address}}).Decode(&profile)

	// TODO: check for no doc err
	if err != nil {
		return entities.Profile{}, err
	}

	return profile, nil
}

func (repo *ProfileRepository) SaveProfile(ctx context.Context, profile entities.Profile) error {
	//TODO: Handle result
	_, err := repo.profiles.UpdateOne(ctx, bson.D{{"_id", profile.Address}}, bson.D{{"$set", profile}}, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}
