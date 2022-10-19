package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
	"github.com/etheralley/etheralley-backend/profiles-api/gateways"
)

func NewGetFungibleToken(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	cacheGateway gateways.ICacheGateway,
	offchainGateway gateways.IOffchainGateway,
) IGetFungibleTokenUseCase {
	return &getFungibleTokenUseCase{
		logger,
		blockchainGateway,
		cacheGateway,
		offchainGateway,
	}
}

type getFungibleTokenUseCase struct {
	logger            common.ILogger
	blockchainGateway gateways.IBlockchainGateway
	cacheGateway      gateways.ICacheGateway
	offchainGateway   gateways.IOffchainGateway
}

type IGetFungibleTokenUseCase interface {
	// Get the metadata and balance of an nft
	Do(ctx context.Context, input *GetFungibleTokenInput) (*entities.FungibleToken, error)
}

type GetFungibleTokenInput struct {
	Address string              `validate:"required,eth_addr"`
	Token   *FungibleTokenInput `validate:"required,dive"`
}

// Fetch balance, name, symbol and decimals concurrently.
// A known token list is maintained in memory that we check first.
// If we miss there, we check the cache.
// If we miss there, we go on chain.
// Name, symbol and decimals are optional implementations and thus we do not bubble an err if we fail to resolve any of them.
func (uc *getFungibleTokenUseCase) Do(ctx context.Context, input *GetFungibleTokenInput) (*entities.FungibleToken, error) {
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
		bal, err := uc.blockchainGateway.GetERC20Balance(ctx, address, contract)

		if err != nil {
			uc.logger.Info(ctx).Err(err).Msgf("err finding token balance: contract address %v blockchain %v", contract.Address, contract.Blockchain)
			return
		}

		balance = &bal
	}()

	go func() {
		defer wg.Done()

		metadata, err := uc.offchainGateway.GetFungibleMetadata(ctx, contract)

		if err == nil {
			uc.logger.Debug(ctx).Msgf("memory hit for token metadata: contract address %v blockchain %v", contract.Address, contract.Blockchain)
			name = metadata.Name
			symbol = metadata.Symbol
			decimals = metadata.Decimals
			return
		}

		uc.logger.Debug(ctx).Msgf("memory miss for token metadata: contract address %v blockchain %v", contract.Address, contract.Blockchain)

		metadata, err = uc.cacheGateway.GetFungibleMetadata(ctx, contract)

		if err == nil {
			uc.logger.Debug(ctx).Msgf("cache hit for token metadata: contract address %v blockchain %v", contract.Address, contract.Blockchain)
			name = metadata.Name
			symbol = metadata.Symbol
			decimals = metadata.Decimals
			return
		}

		uc.logger.Debug(ctx).Msgf("cache miss for token metadata: contract address %v blockchain %v", contract.Address, contract.Blockchain)

		var innerWg sync.WaitGroup
		var nameErr error
		var symbolErr error
		var decimalErr error
		innerWg.Add(3)

		go func() {
			defer innerWg.Done()

			result, err := uc.blockchainGateway.GetERC20Name(ctx, contract)

			if err != nil {
				uc.logger.Info(ctx).Err(err).Msgf("err finding token name: contract address %v blockchain %v", contract.Address, contract.Blockchain)
				nameErr = err
				return
			}

			name = &result
		}()

		go func() {
			defer innerWg.Done()

			result, err := uc.blockchainGateway.GetERC20Symbol(ctx, contract)

			if err != nil {
				uc.logger.Info(ctx).Err(err).Msgf("err finding token symbol: contract address %v blockchain %v", contract.Address, contract.Blockchain)
				symbolErr = err
				return
			}

			symbol = &result
		}()

		go func() {
			defer innerWg.Done()

			result, err := uc.blockchainGateway.GetERC20Decimals(ctx, contract)

			if err != nil {
				uc.logger.Info(ctx).Err(err).Msgf("err finding token decimals: contract address %v blockchain %v", contract.Address, contract.Blockchain)
				decimalErr = err
				return
			}

			decimals = &result
		}()

		innerWg.Wait()

		if nameErr != nil || symbolErr != nil || decimalErr != nil {
			uc.logger.Debug(ctx).Msgf("skipping token metadata cache: contract address %v blockchain %v name err %v symbol err %v decimal err %v", contract.Address, contract.Blockchain, nameErr, symbolErr, decimalErr)
			return
		}

		uc.cacheGateway.SaveFungibleMetadata(ctx, contract, &entities.FungibleMetadata{
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
