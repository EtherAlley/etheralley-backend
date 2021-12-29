package mongo

import (
	"context"
	"fmt"
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

	mongoURI := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v?retryWrites=true&w=majority", settings.MongoUsername, settings.MongoPassword, settings.MongoHost, settings.MongoPort, settings.MongoAdminDB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		logger.Err(err, "mongo connection error")
		panic(err)
	}

	db := client.Database(settings.MongoDB)
	profiles := db.Collection("profiles")

	return &Gateway{
		logger,
		profiles,
	}
}
