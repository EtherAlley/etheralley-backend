package usecases

import (
	"context"
	"errors"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
)

func NewGetNFTUseCase(logger *common.Logger, blockchainGateway *ethereum.Gateway, cacheGateway *redis.Gateway) GetNFTUseCase {
	return GetNFT(logger, blockchainGateway, cacheGateway)
}

// TODO: validate inputs
// concurrent calls to get metadata & validate owner
// metadata doesnt change so we cache it
func GetNFT(logger *common.Logger, blockchainGateway gateways.IBlockchainGateway, cacheGateway gateways.ICacheGateway) GetNFTUseCase {
	return func(ctx context.Context, address string, nftLocation *entities.NFTLocation) (*entities.NFT, error) {
		logger.Debugf("get nft usecase: %v %+v", address, nftLocation)

		var wg sync.WaitGroup
		var metadata *entities.NFTMetadata
		var metadataErr error
		var owned bool
		var ownedErr error

		wg.Add(2)

		go func() {
			defer wg.Done()
			metadata, metadataErr = cacheGateway.GetNFTMetadata(ctx, nftLocation)

			if metadataErr == nil {
				return
			}

			metadata, metadataErr = blockchainGateway.GetNFTMetadata(nftLocation)

			cacheGateway.SaveNFTMetadata(ctx, nftLocation, metadata)
		}()

		go func() {
			defer wg.Done()
			owned, ownedErr = blockchainGateway.VerifyOwner(address, nftLocation)
		}()

		wg.Wait()

		if metadataErr != nil || ownedErr != nil {
			logger.Errorf("get nft usecase: %v %v", metadataErr, ownedErr)
			return nil, errors.New("")
		}

		nft := &entities.NFT{
			Location: nftLocation,
			Owned:    owned,
			Metadata: metadata,
		}

		return nft, nil
	}
}
