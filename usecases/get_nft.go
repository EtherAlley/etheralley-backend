package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type GetNonFungibleTokenInput struct {
	Address          string                 `validate:"required,eth_addr"`
	NonFungibleToken *NonFungibleTokenInput `validate:"required,dive"`
}

// Get the metadata and balance of an nft
//
// Metadata doesnt change so we cache it
//
// Metadata is an optional implementation in ERC721 and ERC1155 and may not exist.
// Its also possible that we simply have issues following the uri.
// In these scenarios we will return nil metadata and not bubble up an err
type IGetNonFungibleTokenUseCase func(ctx context.Context, input *GetNonFungibleTokenInput) (*entities.NonFungibleToken, error)

func NewGetNonFungibleToken(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	offchainGateway gateways.IOffchainGateway,
	cacheGateway gateways.ICacheGateway,
) IGetNonFungibleTokenUseCase {
	return func(ctx context.Context, input *GetNonFungibleTokenInput) (*entities.NonFungibleToken, error) {
		if err := common.ValidateStruct(input); err != nil {
			return nil, err
		}

		address := input.Address
		tokenId := input.NonFungibleToken.TokenId
		contract := &entities.Contract{
			Blockchain: input.NonFungibleToken.Contract.Blockchain,
			Address:    input.NonFungibleToken.Contract.Address,
			Interface:  input.NonFungibleToken.Contract.Interface,
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

			var uri string
			uri, metadataErr = blockchainGateway.GetNonFungibleURI(ctx, contract, tokenId)

			if metadataErr != nil {
				logger.Debugf(ctx, "err getting nft uri: contract address %v token id %v chain %v err %v", contract.Address, tokenId, contract.Blockchain, metadataErr)
				return
			}

			metadata, metadataErr = offchainGateway.GetNonFungibleMetadata(ctx, uri)

			if metadataErr != nil {
				logger.Debugf(ctx, "err getting nft metadata: contract address %v token id %v chain %v err %v", contract.Address, tokenId, contract.Blockchain, metadataErr)
				return
			}

			logger.Debugf(ctx, "found nft metadata: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)

			cacheGateway.SaveNonFungibleMetadata(ctx, contract, tokenId, metadata)
		}()

		go func() {
			defer wg.Done()
			balance, balanceErr = blockchainGateway.GetNonFungibleBalance(ctx, address, contract, tokenId)

			if balanceErr != nil {
				logger.Errf(ctx, balanceErr, "err getting nft balance: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)
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
