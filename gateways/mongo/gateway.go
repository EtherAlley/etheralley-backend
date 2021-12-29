package mongo

import (
	"context"
	"time"

	"github.com/etheralley/etheralley-core-api/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Gateway struct {
	logger   *common.Logger
	profiles *mongo.Collection
}

func NewGateway(settings *common.Settings, logger *common.Logger) *Gateway {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(settings.MongoURI))

	if err != nil {
		panic(err)
	}

	db := client.Database(settings.MongoDB)
	profiles := db.Collection("profiles")

	return &Gateway{
		logger,
		profiles,
	}
}
