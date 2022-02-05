package mongo

import (
	"context"
	"fmt"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Gateway struct {
	logger   common.ILogger
	profiles *mongo.Collection
}

func NewGateway(settings common.ISettings, logger common.ILogger) gateways.IDatabaseGateway {
	ctx := context.Background()

	mongoURI := fmt.Sprintf("mongodb://%v?retryWrites=true&w=majority", settings.DatabaseURI())

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		logger.Err(ctx, err, "mongo connection error")
		panic(err)
	}

	db := client.Database(settings.Database())
	profiles := db.Collection("profiles")

	return &Gateway{
		logger,
		profiles,
	}
}
