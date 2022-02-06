package thegraph

import (
	"context"

	"github.com/etheralley/etheralley-core-api/entities"
)

type SwapsQuery = struct {
	Swaps []struct {
		Timestamp   string `graphql:"timestamp" json:"timestamp"`
		AmountUSD   string `graphql:"amountUSD" json:"amountUSD"`
		Transaction struct {
			Id    string `graphql:"id" json:"id"`
			Swaps []struct {
				Amount0In  string `graphql:"amount0In" json:"amount0In"`
				Amount0Out string `graphql:"amount0Out" json:"amount0Out"`
				Amount1In  string `graphql:"amount1In" json:"amount1In"`
				Amount1Out string `graphql:"amount1Out" json:"amount1Out"`
				Pair       struct {
					Token0 struct {
						Id     string `graphql:"id" json:"id"`
						Symbol string `graphql:"symbol" json:"symbol"`
					} `graphql:"token0" json:"token0"`
					Token1 struct {
						Id     string `graphql:"id" json:"id"`
						Symbol string `graphql:"symbol" json:"symbol"`
					} `graphql:"token1" json:"token1"`
				} `graphql:"pair" json:"pair"`
			} `graphql:"swaps" json:"swaps"`
		} `graphql:"transaction" json:"transaction"`
	} `graphql:"swaps(first: 5, orderBy: amountUSD, orderDirection: desc, where: { to: $to })" json:"swaps"`
}

func (gw *Gateway) GetSwaps(ctx context.Context, address string, contract *entities.Contract) (*[]entities.Swap, error) {
	url, err := gw.GetSubgraphUrl(contract.Blockchain, contract.Interface)

	if err != nil {
		gw.logger.Err(ctx, err, "error building subgraph url")
		return nil, err
	}

	query := &SwapsQuery{}
	variables := map[string]interface{}{
		"to": address,
	}
	err = gw.graphClient.Query(ctx, url, query, variables)

	if err != nil {
		gw.logger.Err(ctx, err, "error calling subgraph")
		return nil, err
	}

	swaps := []entities.Swap{}
	for _, swap := range query.Swaps {
		swapLen := len(swap.Transaction.Swaps)
		if swapLen == 0 {
			continue
		}

		input := entities.SwapToken{}
		output := entities.SwapToken{}

		inputSwap := swap.Transaction.Swaps[0]
		outputSwap := swap.Transaction.Swaps[swapLen-1]

		if inputSwap.Amount0In != "0" {
			input.Amount = inputSwap.Amount0In
			input.Id = inputSwap.Pair.Token0.Id
			input.Symbol = inputSwap.Pair.Token0.Symbol
		} else {
			input.Amount = inputSwap.Amount1In
			input.Id = inputSwap.Pair.Token1.Id
			input.Symbol = inputSwap.Pair.Token1.Symbol
		}

		if outputSwap.Amount0Out != "0" {
			output.Amount = outputSwap.Amount0Out
			output.Id = outputSwap.Pair.Token0.Id
			output.Symbol = outputSwap.Pair.Token0.Symbol
		} else {
			output.Amount = outputSwap.Amount1Out
			output.Id = outputSwap.Pair.Token1.Id
			output.Symbol = outputSwap.Pair.Token1.Symbol
		}

		swaps = append(swaps, entities.Swap{
			Id:        swap.Transaction.Id,
			Timestamp: swap.Timestamp,
			AmountUSD: swap.AmountUSD,
			Input:     &input,
			Output:    &output,
		})
	}

	return &swaps, nil
}
