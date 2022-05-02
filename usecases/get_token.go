package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type GetFungibleTokenInput struct {
	Address string              `validate:"required,eth_addr"`
	Token   *FungibleTokenInput `validate:"required,dive"`
}

// get the metadata and balance of an nft
type IGetFungibleTokenUseCase func(ctx context.Context, input *GetFungibleTokenInput) (*entities.FungibleToken, error)

// fetch balance, name, symbol and decimals concurrently
// a known token list is maintained in memory that we check first
// if we miss there, we check the cache
// if we miss there, we go on chain
// name, symbol and decimals are optional implementations and thus we do not bubble an err if we fail to resolve any of them
func NewGetFungibleToken(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	cacheGateway gateways.ICacheGateway,
	offchainGateway gateways.IOffchainGateway,
) IGetFungibleTokenUseCase {
	return func(ctx context.Context, input *GetFungibleTokenInput) (*entities.FungibleToken, error) {
		if err := common.ValidateStruct(input); err != nil {
			return nil, err
		}

		address := input.Address
		contract := &entities.Contract{
			Blockchain: input.Token.Contract.Blockchain,
			Address:    input.Token.Contract.Address,
			Interface:  input.Token.Contract.Interface,
		}

		var balance *string
		var name *string
		var symbol *string
		var decimals *uint8

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			bal, err := blockchainGateway.GetERC20Balance(ctx, address, contract)

			if err != nil {
				logger.Errf(ctx, err, "err finding token balance: contract address %v blockchain %v", contract.Address, contract.Blockchain)
				return
			}

			balance = &bal
		}()

		go func() {
			defer wg.Done()

			metadata, err := offchainGateway.GetFungibleMetadata(ctx, contract)

			if err == nil {
				logger.Debugf(ctx, "memory hit for token metadata: contract address %v blockchain %v", contract.Address, contract.Blockchain)
				name = metadata.Name
				symbol = metadata.Symbol
				decimals = metadata.Decimals
				return
			}

			logger.Debugf(ctx, "memory miss for token metadata: contract address %v blockchain %v", contract.Address, contract.Blockchain)

			metadata, err = cacheGateway.GetFungibleMetadata(ctx, contract)

			if err == nil {
				logger.Debugf(ctx, "cache hit for token metadata: contract address %v blockchain %v", contract.Address, contract.Blockchain)
				name = metadata.Name
				symbol = metadata.Symbol
				decimals = metadata.Decimals
				return
			}

			logger.Debugf(ctx, "cache miss for token metadata: contract address %v blockchain %v", contract.Address, contract.Blockchain)

			var innerWg sync.WaitGroup
			var nameErr error
			var symbolErr error
			var decimalErr error
			innerWg.Add(3)

			go func() {
				defer innerWg.Done()

				result, err := blockchainGateway.GetERC20Name(ctx, contract)

				if err != nil {
					logger.Errf(ctx, err, "err finding token name: contract address %v blockchain %v", contract.Address, contract.Blockchain)
					nameErr = err
					return
				}

				name = &result
			}()

			go func() {
				defer innerWg.Done()

				result, err := blockchainGateway.GetERC20Symbol(ctx, contract)

				if err != nil {
					logger.Errf(ctx, err, "err finding token symbol: contract address %v blockchain %v", contract.Address, contract.Blockchain)
					symbolErr = err
					return
				}

				symbol = &result
			}()

			go func() {
				defer innerWg.Done()

				result, err := blockchainGateway.GetERC20Decimals(ctx, contract)

				if err != nil {
					logger.Errf(ctx, err, "err finding token decimals: contract address %v blockchain %v", contract.Address, contract.Blockchain)
					decimalErr = err
					return
				}

				decimals = &result
			}()

			innerWg.Wait()

			if nameErr != nil || symbolErr != nil || decimalErr != nil {
				logger.Debugf(ctx, "skipping token metadata cache: contract address %v blockchain %v", contract.Address, contract.Blockchain)
				return
			}

			cacheGateway.SaveFungibleMetadata(ctx, contract, &entities.FungibleMetadata{
				Name:     name,
				Symbol:   symbol,
				Decimals: decimals,
			})
		}()

		wg.Wait()

		return &entities.FungibleToken{
			Contract: contract,
			Balance:  balance,
			Metadata: &entities.FungibleMetadata{
				Name:     name,
				Symbol:   symbol,
				Decimals: decimals,
			},
		}, nil
	}
}
