package mongo

import (
	"context"

	"github.com/eflem00/go-example-app/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDb(settings *common.Settings, lgr *common.Logger) *mongo.Database {
	//TODO: Fix this context
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(settings.MongoDBURI))

	if err != nil {
		panic(err)
	}

	return client.Database("etheralley")
}
