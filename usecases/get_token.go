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
// cache metadata
// name, symbol and decimals are optional implementations and thus we do not bubble an err if we fail to resolve any of them
func NewGetFungibleToken(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	cacheGateway gateways.ICacheGateway,
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
			bal, err := blockchainGateway.GetFungibleBalance(ctx, address, contract)

			if err != nil {
				logger.Errf(ctx, err, "err finding token balance: contract address %v blockchain %v", contract.Address, contract.Blockchain)
				return
			}

			balance = &bal
		}()

		go func() {
			defer wg.Done()

			metadata, err := cacheGateway.GetFungibleMetadata(ctx, contract)

			if err == nil {
				logger.Debugf(ctx, "cache hit for token metadata: contract address %v blockchain %v", contract.Address, contract.Blockchain)
				name = metadata.Name
				symbol = metadata.Symbol
				decimals = metadata.Decimals
				return
			}

			logger.Debugf(ctx, "cache miss for token metadata: contract address %v blockchain %v", contract.Address, contract.Blockchain)

			var innerWg sync.WaitGroup
			var metadataErr error // err is written to by all three goroutines. if any of them have an err we will not cache the token metadata
			innerWg.Add(3)

			go func() {
				defer innerWg.Done()

				nameResult, nameErr := blockchainGateway.GetFungibleName(ctx, contract)

				if nameErr != nil {
					logger.Errf(ctx, nameErr, "err finding token name: contract address %v blockchain %v", contract.Address, contract.Blockchain)
					metadataErr = nameErr
					return
				}

				name = &nameResult
			}()

			go func() {
				defer innerWg.Done()

				symb, symbolErr := blockchainGateway.GetFungibleSymbol(ctx, contract)

				if symbolErr != nil {
					logger.Errf(ctx, symbolErr, "err finding token symbol: contract address %v blockchain %v", contract.Address, contract.Blockchain)
					metadataErr = symbolErr
					return
				}

				symbol = &symb
			}()

			go func() {
				defer innerWg.Done()

				dec, decimalsErr := blockchainGateway.GetFungibleDecimals(ctx, contract)

				if decimalsErr != nil {
					logger.Errf(ctx, decimalsErr, "err finding token decimals: contract address %v blockchain %v", contract.Address, contract.Blockchain)
					metadataErr = decimalsErr
					return
				}

				decimals = &dec
			}()

			innerWg.Wait()

			if metadataErr != nil {
				logger.Debugf(ctx, "skipping token metadata cache: contract address %v blockchain %v err %v", contract.Address, contract.Blockchain, metadataErr)
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
