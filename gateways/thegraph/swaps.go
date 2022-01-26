package thegraph

import (
	"context"

	"github.com/etheralley/etheralley-core-api/entities"
)

type SwapsQuery = struct {
	Swaps []struct {
		Id         string `json:"id"`
		Timestamp  string `json:"timestamp"`
		AmountUSD  string `graphql:"amountUSD" json:"amountUSD"`
		Amount0In  string `graphql:"amount0In" json:"amount0In"`
		Amount0Out string `graphql:"amount0Out" json:"amount0Out"`
		Amount1In  string `graphql:"amount1In" json:"amount1In"`
		Amount1Out string `graphql:"amount1Out" json:"amount1Out"`
		Pair       struct {
			Token0 struct {
				Id     string `json:"id"`
				Symbol string `json:"symbol"`
			} `json:"token0"`
			Token1 struct {
				Id     string `json:"id"`
				Symbol string `json:"symbol"`
			} `json:"token1"`
		} `json:"pair"`
	} `graphql:"swaps(first: 5, orderBy: amountUSD, orderDirection: desc, where: { sender: $sender })" json:"swaps"`
}

func (gw *Gateway) GetSwaps(ctx context.Context, address string, contract *entities.Contract) (entities.StatisticalData, error) {
	url, err := gw.GetSubgraphUrl(contract.Blockchain, contract.Interface)

	if err != nil {
		gw.logger.Err(err, "error building subgraph url")
		return nil, err
	}

	query := &SwapsQuery{}
	variables := map[string]interface{}{
		"sender": address,
	}
	err = gw.client.Query(ctx, url, query, variables)

	if err != nil {
		gw.logger.Err(err, "error calling subgraph")
		return nil, err
	}

	return query, nil
}
