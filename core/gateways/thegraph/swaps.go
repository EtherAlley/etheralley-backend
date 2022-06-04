package thegraph

import (
	"context"
	"fmt"
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/core/entities"
)

type swapJson = struct {
	Id        string         `json:"id"`
	Timestamp string         `json:"timestamp"`
	AmountUSD string         `json:"amountUSD"`
	Input     *swapTokenJson `json:"input"`
	Output    *swapTokenJson `json:"output"`
}

type swapTokenJson = struct {
	Id     string `json:"id"`
	Amount string `json:"amount"`
	Symbol string `json:"symbol"`
}

// in general, regardless of the subgraph schema, we gets swaps where the address is there recipient and get the full swaps from the associated transaction
// we use the full swap list to reconstruct the full swap. We do this because in many scenarios there are multiple hops a swap takes to get to its final destination
// TODO:
// https://thegraph.com/hosted-service/subgraph/ianlapham/uniswap-arbitrum-one
// https://thegraph.com/hosted-service/subgraph/ianlapham/uniswap-optimism
func (gw *gateway) GetSwaps(ctx context.Context, address string, contract *entities.Contract) (interface{}, error) {
	hostedURI := gw.settings.TheGraphHostedURI()

	switch contract.Blockchain {
	case common.ETHEREUM:
		switch contract.Interface {
		case common.SUSHISWAP_EXCHANGE:
			return gw.getSwapsSushiAndUniV2(ctx, fmt.Sprintf("%v/sushiswap/exchange", hostedURI), address)
		case common.UNISWAP_V2_EXCHANGE:
			return gw.getSwapsSushiAndUniV2(ctx, fmt.Sprintf("%v/uniswap/uniswap-v2", hostedURI), address)
		case common.UNISWAP_V3_EXCHANGE:
			return gw.getSwapsUniV3(ctx, fmt.Sprintf("%v/uniswap/uniswap-v3", hostedURI), address)
		}
	case common.POLYGON:
		switch contract.Interface {
		case common.SUSHISWAP_EXCHANGE:
			return gw.getSwapsSushiAndUniV2(ctx, fmt.Sprintf("%v/sushiswap/matic-exchange", hostedURI), address)
		case common.UNISWAP_V3_EXCHANGE:
			return gw.getSwapsUniV3(ctx, fmt.Sprintf("%v/ianlapham/uniswap-v3-polygon", hostedURI), address)
		}
	case common.ARBITRUM:
		switch contract.Interface {
		case common.SUSHISWAP_EXCHANGE:
			return gw.getSwapsSushiAndUniV2(ctx, fmt.Sprintf("%v/sushiswap/arbitrum-exchange", hostedURI), address)
		}
	}

	return nil, fmt.Errorf("invalid subgraph for blockchain %v and interface %v", contract.Blockchain, contract.Interface)
}

type sushiAndUniV2SwapsQuery = struct {
	Swaps []struct {
		Timestamp  string `graphql:"timestamp" json:"timestamp"`
		AmountUSD  string `graphql:"amountUSD" json:"amountUSD"`
		To         string `graphql:"to" json:"to"`
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
		Transaction struct {
			Id    string `graphql:"id" json:"id"`
			Swaps []struct {
				To         string `graphql:"to" json:"to"`
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

func (gw *gateway) getSwapsSushiAndUniV2(ctx context.Context, url string, address string) (*[]swapJson, error) {
	query := &sushiAndUniV2SwapsQuery{}
	variables := map[string]interface{}{
		"to": address,
	}
	err := gw.graphClient.Query(ctx, url, query, variables)

	if err != nil {
		return nil, fmt.Errorf("getSwapsSushiAndUniV2 query %w", err)
	}

	swaps := []swapJson{}
	for _, swap := range query.Swaps {
		swapLen := len(swap.Transaction.Swaps)
		if swapLen == 0 {
			continue
		}

		var input swapTokenJson
		var output swapTokenJson

		if swap.Amount0In != "0" {
			input = swapTokenJson{
				Amount: swap.Amount0In,
				Id:     swap.Pair.Token0.Id,
				Symbol: swap.Pair.Token0.Symbol,
			}
			output = swapTokenJson{
				Amount: swap.Amount1Out,
				Id:     swap.Pair.Token1.Id,
				Symbol: swap.Pair.Token1.Symbol,
			}
		} else {
			input = swapTokenJson{
				Amount: swap.Amount1In,
				Id:     swap.Pair.Token1.Id,
				Symbol: swap.Pair.Token1.Symbol,
			}
			output = swapTokenJson{
				Amount: swap.Amount0Out,
				Id:     swap.Pair.Token0.Id,
				Symbol: swap.Pair.Token0.Symbol,
			}
		}

		// TODO: This only finds one other link in the chain of swaps when trying to find the original input token.
		for _, tSwap := range swap.Transaction.Swaps {
			if tSwap.To != swap.To {
				if tSwap.Amount0Out == input.Amount && tSwap.Pair.Token0.Id == input.Id {
					input = swapTokenJson{
						Amount: tSwap.Amount1In,
						Id:     tSwap.Pair.Token1.Id,
						Symbol: tSwap.Pair.Token1.Symbol,
					}
				} else if tSwap.Amount1Out == input.Amount && tSwap.Pair.Token1.Id == input.Id {
					input = swapTokenJson{
						Amount: tSwap.Amount0In,
						Id:     tSwap.Pair.Token0.Id,
						Symbol: tSwap.Pair.Token0.Symbol,
					}
				}
			}
		}

		swaps = append(swaps, swapJson{
			Id:        swap.Transaction.Id,
			Timestamp: swap.Timestamp,
			AmountUSD: swap.AmountUSD,
			Input:     &input,
			Output:    &output,
		})
	}

	return &swaps, nil
}

type uniV3SwapsQuery = struct {
	Swaps []struct {
		Timestamp string `graphql:"timestamp" json:"timestamp"`
		AmountUSD string `graphql:"amountUSD" json:"amountUSD"`
		Recipient string `graphql:"recipient" json:"recipient"`
		Amount0   string `graphql:"amount0" json:"amount0"`
		Amount1   string `graphql:"amount1" json:"amount1"`
		Token0    struct {
			Id     string `graphql:"id" json:"id"`
			Symbol string `graphql:"symbol" json:"symbol"`
		} `graphql:"token0" json:"token0"`
		Token1 struct {
			Id     string `graphql:"id" json:"id"`
			Symbol string `graphql:"symbol" json:"symbol"`
		} `graphql:"token1" json:"token1"`
		Transaction struct {
			Id    string `graphql:"id" json:"id"`
			Swaps []struct {
				Recipient string `graphql:"recipient" json:"recipient"`
				Amount0   string `graphql:"amount0" json:"amount0"`
				Amount1   string `graphql:"amount1" json:"amount1"`
				Token0    struct {
					Id     string `graphql:"id" json:"id"`
					Symbol string `graphql:"symbol" json:"symbol"`
				} `graphql:"token0" json:"token0"`
				Token1 struct {
					Id     string `graphql:"id" json:"id"`
					Symbol string `graphql:"symbol" json:"symbol"`
				} `graphql:"token1" json:"token1"`
			} `graphql:"swaps" json:"swaps"`
		} `graphql:"transaction" json:"transaction"`
	} `graphql:"swaps(first: 10, orderBy: amountUSD, orderDirection: desc, where: { recipient: $recipient })" json:"swaps"`
}

func (gw *gateway) getSwapsUniV3(ctx context.Context, url string, address string) (*[]swapJson, error) {
	query := &uniV3SwapsQuery{}
	variables := map[string]interface{}{
		"recipient": address,
	}
	err := gw.graphClient.Query(ctx, url, query, variables)

	if err != nil {
		return nil, fmt.Errorf("getSwapsUniV3 query %w", err)
	}

	swaps := []swapJson{}
	for _, swap := range query.Swaps {
		swapLen := len(swap.Transaction.Swaps)
		if swapLen == 0 {
			continue
		}

		var input swapTokenJson
		var output swapTokenJson

		if strings.Contains(swap.Amount0, "-") {
			output = swapTokenJson{
				Amount: strings.Replace(swap.Amount0, "-", "", 1),
				Id:     swap.Token0.Id,
				Symbol: swap.Token0.Symbol,
			}
			input = swapTokenJson{
				Amount: swap.Amount1,
				Id:     swap.Token1.Id,
				Symbol: swap.Token1.Symbol,
			}
		} else if strings.Contains(swap.Amount1, "-") {
			output = swapTokenJson{
				Amount: strings.Replace(swap.Amount1, "-", "", 1),
				Id:     swap.Token1.Id,
				Symbol: swap.Token1.Symbol,
			}
			input = swapTokenJson{
				Amount: swap.Amount0,
				Id:     swap.Token0.Id,
				Symbol: swap.Token0.Symbol,
			}
		}

		// TODO: This only finds one other link in the chain of swaps when trying to find the original input token.
		for _, tSwap := range swap.Transaction.Swaps {
			amount0 := strings.Replace(tSwap.Amount0, "-", "", 1)
			amount1 := strings.Replace(tSwap.Amount1, "-", "", 1)
			if tSwap.Recipient != swap.Recipient {
				if amount0 == input.Amount && tSwap.Token0.Id == input.Id {
					input = swapTokenJson{
						Amount: amount1,
						Id:     tSwap.Token1.Id,
						Symbol: tSwap.Token1.Symbol,
					}
				} else if amount1 == input.Amount && tSwap.Token1.Id == input.Id {
					input = swapTokenJson{
						Amount: amount0,
						Id:     tSwap.Token0.Id,
						Symbol: tSwap.Token0.Symbol,
					}
				}
			}
		}

		swaps = append(swaps, swapJson{
			Id:        swap.Transaction.Id,
			Timestamp: swap.Timestamp,
			AmountUSD: swap.AmountUSD,
			Input:     &input,
			Output:    &output,
		})
	}

	return &swaps, nil
}
