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
		var balance *string

		wg.Add(2)

		go func() {
			defer wg.Done()
			mdata, err := cacheGateway.GetNonFungibleMetadata(ctx, contract, tokenId)

			if err == nil {
				logger.Debugf(ctx, "cache hit for nft metadata: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)
				metadata = mdata
				return
			}

			logger.Debugf(ctx, "cache miss for nft metadata: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)

			var uri string
			uri, err = blockchainGateway.GetNonFungibleURI(ctx, contract, tokenId)

			if err != nil {
				logger.Errf(ctx, err, "err getting nft uri: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)
				metadata = nil
				return
			}

			mdata, err = offchainGateway.GetNonFungibleMetadata(ctx, uri)

			if err != nil {
				logger.Errf(ctx, err, "err getting nft metadata: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)
				metadata = nil
				return
			}

			cacheGateway.SaveNonFungibleMetadata(ctx, contract, tokenId, mdata)

			metadata = mdata
		}()

		go func() {
			defer wg.Done()
			bal, err := blockchainGateway.GetNonFungibleBalance(ctx, address, contract, tokenId)

			if err != nil {
				logger.Errf(ctx, err, "err getting nft balance: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)
				balance = nil
				return
			}

			balance = &bal
		}()

		wg.Wait()

		nft := &entities.NonFungibleToken{
			Contract: contract,
			TokenId:  tokenId,
			Balance:  balance,
			Metadata: metadata,
		}

		return nft, nil
	}
}
