package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

func NewGetAllNFTsUseCase(logger *common.Logger, getNFTUseCase GetNFTUseCase) GetAllNFTsUseCase {
	return GetAllNFTs(logger, getNFTUseCase)
}

// fetch each hydrated nft concurrently
// we can use a simple slice here since each result in the go routine writes to its own index location
func GetAllNFTs(logger *common.Logger, getNFTUseCase GetNFTUseCase) GetAllNFTsUseCase {
	return func(ctx context.Context, address string, nftLocations *[]entities.NFTLocation) *[]entities.NFT {
		var wg sync.WaitGroup

		nfts := make([]*entities.NFT, len(*nftLocations))

		for i, nftLocation := range *nftLocations {
			wg.Add(1)

			go func(i int, loc entities.NFTLocation) {
				defer wg.Done()

				nft, err := getNFTUseCase(ctx, address, &loc)

				if err == nil && nft.Owned {
					nfts[i] = nft
				} else {
					logger.Debugf("invalid nft provided: %v", err)
				}

			}(i, nftLocation)
		}

		wg.Wait()

		trimmedNfts := []entities.NFT{}
		for _, nft := range nfts {
			if nft != nil {
				trimmedNfts = append(trimmedNfts, *nft)
			}
		}

		return &trimmedNfts
	}
}
