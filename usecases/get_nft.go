package usecases

import (
	"context"
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
				logger.Debugf("cache hit for nft: contract address %v token id %v chain %v", nftLocation.ContractAddress, nftLocation.TokenId, nftLocation.Blockchain)
				return
			}

			logger.Debugf("cache miss for nft: contract address %v token id %v chain %v", nftLocation.ContractAddress, nftLocation.TokenId, nftLocation.Blockchain)

			metadata, metadataErr = blockchainGateway.GetNFTMetadata(nftLocation)

			if metadataErr != nil {
				logger.Errf(metadataErr, "chain miss for nft: contract address %v token id %v chain %v", nftLocation.ContractAddress, nftLocation.TokenId, nftLocation.Blockchain)
				return
			}

			logger.Debugf("chain hit for nft: contract address %v token id %v chain %v", nftLocation.ContractAddress, nftLocation.TokenId, nftLocation.Blockchain)

			cacheGateway.SaveNFTMetadata(ctx, nftLocation, metadata)
		}()

		go func() {
			defer wg.Done()
			owned, ownedErr = blockchainGateway.VerifyOwner(address, nftLocation)

			if ownedErr != nil {
				logger.Errf(ownedErr, "err verifying owner on-chain for nft: contract address %v token id %v chain %v", nftLocation.ContractAddress, nftLocation.TokenId, nftLocation.Blockchain)
			}
		}()

		wg.Wait()

		if metadataErr != nil {
			return nil, metadataErr
		}

		if ownedErr != nil {
			return nil, ownedErr
		}

		nft := &entities.NFT{
			Location: nftLocation,
			Owned:    owned,
			Metadata: metadata,
		}

		return nft, nil
	}
}
