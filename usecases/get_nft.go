package usecases

import (
	"context"
	"errors"
	"fmt"
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
	settings common.ISettings,
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
			switch contract.Interface {
			case common.ERC721:
				uri, err = blockchainGateway.GetERC721URI(ctx, contract, tokenId)
			case common.ERC1155:
				uri, err = blockchainGateway.GetERC1155URI(ctx, contract, tokenId)
			case common.ENS_REGISTRAR:
				uri, err = fmt.Sprintf("%v/%v/%v", settings.ENSMetadataURI(), contract.Address, tokenId), nil
			case common.CRYPTO_PUNKS:
				uri, err = "", nil
			case common.CRYPTO_KITTIES:
				uri, err = "", nil
			default:
				uri, err = "", errors.New("invalida schema name")
			}

			if err != nil {
				logger.Errf(ctx, err, "err getting nft uri: contract address %v token id %v chain %v", contract.Address, tokenId, contract.Blockchain)
				metadata = nil
				return
			}

			switch contract.Interface {
			case common.ERC721:
				mdata, err = offchainGateway.GetNonFungibleMetadata(ctx, uri)
			case common.ERC1155:
				mdata, err = offchainGateway.GetNonFungibleMetadata(ctx, uri)
			case common.ENS_REGISTRAR:
				mdata, err = offchainGateway.GetNonFungibleMetadata(ctx, uri)
			case common.CRYPTO_PUNKS:
				mdata, err = offchainGateway.GetPunkMetadata(ctx, tokenId)
			case common.CRYPTO_KITTIES:
				mdata, err = offchainGateway.GetKittieMetadata(ctx, tokenId)
			default:
				mdata, err = nil, errors.New("invalida schema name")
			}

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

			var bal string
			var err error
			switch contract.Interface {
			case common.ERC721:
				bal, err = blockchainGateway.GetERC721Balance(ctx, address, contract, tokenId)
			case common.ERC1155:
				bal, err = blockchainGateway.GetERC1155Balance(ctx, address, contract, tokenId)
			case common.ENS_REGISTRAR:
				bal, err = blockchainGateway.GetERC721Balance(ctx, address, contract, tokenId)
			case common.CRYPTO_PUNKS:
				bal, err = blockchainGateway.GetPunkBalance(ctx, address, contract, tokenId)
			case common.CRYPTO_KITTIES:
				bal, err = blockchainGateway.GetERC721Balance(ctx, address, contract, tokenId)
			default:
				bal, err = "", errors.New("invalida schema name")
			}

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
