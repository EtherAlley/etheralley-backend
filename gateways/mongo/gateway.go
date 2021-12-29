package mongo

import (
	"context"

	"github.com/eflem00/go-example-app/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Gateway struct {
	logger   *common.Logger
	profiles *mongo.Collection
}

func NewGateway(settings *common.Settings, logger *common.Logger) *Gateway {
	//TODO: Fix this context
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(settings.MongoDBURI))

	if err != nil {
		panic(err)
	}

	db := client.Database("etheralley")
	profiles := db.Collection("profiles")

	return &Gateway{
		logger,
		profiles,
	}
}
