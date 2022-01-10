package usecases

import (
	"context"
	"math/big"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

func NewGetAllNonFungibleTokensUseCase(logger *common.Logger, getNonFungibleToken GetNonFungibleTokenUseCase) GetAllNonFungibleTokensUseCase {
	return GetAllNonFungibleTokens(logger, getNonFungibleToken)
}

// for each partial nft provided, fetch the hydrated nft concurrently
// we can use a simple slice here since each result in the go routine writes to its own index location
// balances that are zero for the given address are dropped
func GetAllNonFungibleTokens(logger *common.Logger, getNonFungibleToken GetNonFungibleTokenUseCase) GetAllNonFungibleTokensUseCase {
	return func(ctx context.Context, address string, partials *[]entities.NonFungibleToken) *[]entities.NonFungibleToken {
		var wg sync.WaitGroup

		nfts := make([]*entities.NonFungibleToken, len(*partials))

		for i, partial := range *partials {
			wg.Add(1)

			go func(i int, p entities.NonFungibleToken) {
				defer wg.Done()

				nft, err := getNonFungibleToken(ctx, address, p.Contract, p.TokenId)

				if err != nil {
					logger.Errf(err, "invalid nft: contract address %v token id %v chain %v", p.Contract.Address, p.TokenId, p.Contract.Blockchain)
					return
				}

				balance := new(big.Int)
				balance.SetString(nft.Balance, 10)
				if balance.Cmp(big.NewInt(0)) == 1 {
					nfts[i] = nft
				}

			}(i, partial)
		}

		wg.Wait()

		trimmedNfts := []entities.NonFungibleToken{}
		for _, nft := range nfts {
			if nft != nil {
				trimmedNfts = append(trimmedNfts, *nft)
			}
		}

		return &trimmedNfts
	}
}
