package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

func NewHydrateNFTsUseCase(logger *common.Logger, getNFTUseCase GetNFTUseCase) HydrateNFTsUseCase {
	return HydrateNFTs(logger, getNFTUseCase)
}

func HydrateNFTs(logger *common.Logger, getNFTUseCase GetNFTUseCase) HydrateNFTsUseCase {
	return func(ctx context.Context, address string, partialNFTs []entities.NFT) []entities.NFT {
		nfts := []entities.NFT{}
		for _, partialNft := range partialNFTs {
			nft, err := getNFTUseCase(ctx, address, partialNft.Blockchain, partialNft.ContractAddress, partialNft.SchemaName, partialNft.TokenId)
			if err == nil && nft.Owned {
				nfts = append(nfts, *nft)
			} else {
				logger.Debugf("invalid nft provided: %v", err)
			}
		}
		return nfts
	}
}
