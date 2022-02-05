package common

import (
	"context"
	"net/http"

	"github.com/hasura/go-graphql-client"
)

type IGraphQLClient interface {
	Query(ctx context.Context, url string, q interface{}, v map[string]interface{}) error
}

type graphQLClient struct {
	logger     ILogger
	httpClient *http.Client
}

func NewGraphQLClient(logger ILogger) IGraphQLClient {
	return &graphQLClient{
		httpClient: &http.Client{},
		logger:     logger,
	}
}

func (g *graphQLClient) Query(ctx context.Context, url string, q interface{}, v map[string]interface{}) error {
	client := graphql.NewClient(url, g.httpClient)

	g.logger.Debugf(ctx, "graphql request %v", url)

	return client.Query(ctx, q, v)
}
