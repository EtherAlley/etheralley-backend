package common

import (
	"context"
	"net/http"
	"time"

	"github.com/hasura/go-graphql-client"
)

type IGraphQLClient interface {
	Query(ctx context.Context, url string, q interface{}, v map[string]interface{}) error
}

type graphQLClient struct {
	httpClient *http.Client
}

func NewGraphQLClient() IGraphQLClient {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}
	return &graphQLClient{
		httpClient,
	}
}

func (g *graphQLClient) Query(ctx context.Context, url string, q interface{}, v map[string]interface{}) error {
	client := graphql.NewClient(url, g.httpClient)

	return client.Query(ctx, q, v)
}
