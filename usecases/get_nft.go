package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

// concurrent calls to get metadata & validate owner
// metadata doesnt change so we cache it
// metadata is an optional implementation in ERC721 and ERC1155 and may not exist
// its also possible that we simply have issues following the uri
// in these scenarios we will return nil metadata
func NewGetNonFungibleToken(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	cacheGateway gateways.ICacheGateway,
) IGetNonFungibleTokenUseCase {
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
				logger.Debugf(ctx, "cache hit for nft metadata: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)
				return
			}

			logger.Debugf(ctx, "cache miss for nft metadata: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)

			metadata, metadataErr = blockchainGateway.GetNonFungibleMetadata(ctx, contract, tokenId)

			if metadataErr != nil {
				logger.Debugf(ctx, "err finding nft metadata: contract address %v token id %v chain %v err %v", contract.Address, tokenId, contract.Blockchain, metadataErr)
				return
			}

			logger.Debugf(ctx, "found nft metadata: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)

			cacheGateway.SaveNonFungibleMetadata(ctx, contract, tokenId, metadata)
		}()

		go func() {
			defer wg.Done()
			balance, balanceErr = blockchainGateway.GetNonFungibleBalance(ctx, address, contract, tokenId)

			if balanceErr != nil {
				logger.Errf(ctx, balanceErr, "err verifying nft ownership: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)
			}
		}()

		wg.Wait()

		if balanceErr != nil {
			return nil, balanceErr
		}

		// getting metadata, particularly following urls is a flakey operation. we are intentionally not bubbling up an error here and simply returning nil metadata.
		// if the contract provided is bad it can be detected and bubbled in the balance error
		if metadataErr != nil {
			metadata = nil
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
