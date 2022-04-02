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

		var balance string
		var name string
		var symbol string
		var decimals uint8
		var balanceErr error

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			balance, balanceErr = blockchainGateway.GetFungibleBalance(ctx, address, contract)
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
			innerWg.Add(3)

			var nameErr error
			var symbolErr error
			var decimalsErr error

			go func() {
				defer innerWg.Done()
				name, nameErr = blockchainGateway.GetFungibleName(ctx, contract)
			}()

			go func() {
				defer innerWg.Done()
				symbol, symbolErr = blockchainGateway.GetFungibleSymbol(ctx, contract)
			}()

			go func() {
				defer innerWg.Done()
				decimals, decimalsErr = blockchainGateway.GetFungibleDecimals(ctx, contract)
			}()

			innerWg.Wait()

			if nameErr != nil || symbolErr != nil || decimalsErr != nil {
				return
			}

			cacheGateway.SaveFungibleMetadata(ctx, contract, &entities.FungibleMetadata{
				Name:     name,
				Symbol:   symbol,
				Decimals: decimals,
			})

		}()

		wg.Wait()

		if balanceErr != nil {
			logger.Debugf(ctx, "err finding token balance: contract address %v blockchain %v err %v", contract.Address, contract.Blockchain, balanceErr)
			return nil, balanceErr
		}

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
