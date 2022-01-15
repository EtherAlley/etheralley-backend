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

func NewGetNonFungibleTokenUseCase(logger *common.Logger, blockchainGateway *ethereum.Gateway, cacheGateway *redis.Gateway) GetNonFungibleTokenUseCase {
	return GetNonFungibleToken(logger, blockchainGateway, cacheGateway)
}

// concurrent calls to get metadata & validate owner
// metadata doesnt change so we cache it
// metadata is an optional implementation in ERC721 and ERC1155 but we don't support presenting an nft with no metadata in the ui
func GetNonFungibleToken(logger *common.Logger, blockchainGateway gateways.IBlockchainGateway, cacheGateway gateways.ICacheGateway) GetNonFungibleTokenUseCase {
	return func(ctx context.Context, address string, contract *entities.Contract, tokenId string) (*entities.NonFungibleToken, error) {
		if err := common.ValidateStruct(contract); err != nil {
			return nil, err
		}

		if err := common.ValidateField(tokenId, `required,numeric`); err != nil {
			return nil, err
		}

		var wg sync.WaitGroup
		var metadata *entities.NonFungibleMetadata
		var metadataErr error
		var balance string
		var balanceErr error

		wg.Add(2)

		go func() {
			defer wg.Done()
			metadata, metadataErr = cacheGateway.GetNonFungibleMetadata(ctx, contract, tokenId)

			if metadataErr == nil {
				logger.Debugf("cache hit for nft metadata: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)
				return
			}

			logger.Debugf("cache miss for nft metadata: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)

			metadata, metadataErr = blockchainGateway.GetNonFungibleMetadata(contract, tokenId)

			if metadataErr != nil {
				logger.Debugf("err finding nft metadata: contract address %v token id %v chain %v err %v", contract.Address, tokenId, contract.Blockchain, metadataErr)
				return
			}

			logger.Debugf("found nft metadata: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)

			cacheGateway.SaveNonFungibleMetadata(ctx, contract, tokenId, metadata)
		}()

		go func() {
			defer wg.Done()
			balance, balanceErr = blockchainGateway.GetNonFungibleBalance(address, contract, tokenId)

			if balanceErr != nil {
				logger.Errf(balanceErr, "err verifying nft ownership: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)
			}
		}()

		wg.Wait()

		if balanceErr != nil {
			return nil, balanceErr
		}

		if metadataErr != nil {
			return nil, metadataErr
		}

		nft := &entities.NonFungibleToken{
			Contract: contract,
			TokenId:  tokenId,
			Balance:  balance,
			Metadata: metadata,
		}

		return nft, nil
	}
}
