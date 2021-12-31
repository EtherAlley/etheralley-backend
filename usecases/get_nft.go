package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum"
)

func NewGetNFTUseCase(logger *common.Logger, gateway *ethereum.Gateway) GetNFTUseCase {
	return GetNFT(logger, gateway)
}

func GetNFT(logger *common.Logger, gateway gateways.IBlockchainGateway) GetNFTUseCase {
	return func(ctx context.Context, address string, blockchain string, contractAddress string, schemaName string, tokenId string) (*entities.NFT, error) {
		logger.Infof("get nft info usecase %v %v %v %v %v", address, blockchain, contractAddress, schemaName, tokenId)

		// TODO: do a bunch of input validation

		// TODO: call gateway based on blockchain and call metho of gateway based on schemaname?

		return &entities.NFT{}, nil
	}
}
