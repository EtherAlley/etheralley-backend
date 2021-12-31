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

// TODO: validate inputs
// TODO: cache metadata
// TODO: concurrent calls to get metadata/validate owner?
func GetNFT(logger *common.Logger, gateway gateways.IBlockchainGateway) GetNFTUseCase {
	return func(ctx context.Context, address string, blockchain string, contractAddress string, schemaName string, tokenId string) (*entities.NFT, error) {
		logger.Debugf("get nft usecase: address: %v blockchain: %v contractAddress: %v schemaName: %v tokenId: %v", address, blockchain, contractAddress, schemaName, tokenId)

		metadata, err := gateway.GetNFTMetadata(contractAddress, tokenId, schemaName)

		if err != nil {
			logger.Err(err, "get nft usecase: ")
			return nil, err
		}

		owner, err := gateway.VerifyOwner(contractAddress, address, tokenId, schemaName)

		if err != nil {
			logger.Err(err, "get nft usecase: ")
			return nil, err
		}

		nft := &entities.NFT{
			TokenId:         tokenId,
			Blockchain:      blockchain,
			ContractAddress: contractAddress,
			SchemaName:      schemaName,
			Owned:           owner,
			Metadata:        *metadata,
		}

		return nft, nil
	}
}
