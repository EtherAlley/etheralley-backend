package common

import (
	"context"
	"net/http"
	"time"

	"github.com/hasura/go-graphql-client"
)

type GraphQLClient struct {
	httpClient *http.Client
}

func NewGraphQLClient() *GraphQLClient {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}
	return &GraphQLClient{
		httpClient,
	}
}

func (g *GraphQLClient) Query(ctx context.Context, url string, q interface{}, v map[string]interface{}) error {
	client := graphql.NewClient(url, g.httpClient)

	return client.Query(ctx, q, v)
}
