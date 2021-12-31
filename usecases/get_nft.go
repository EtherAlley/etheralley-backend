package usecases

import (
	"context"
	"encoding/json"
	"errors"

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
		logger.Infof("GetNFT %v %v %v %v %v", address, blockchain, contractAddress, schemaName, tokenId)

		// TODO: do a bunch of input validation

		// TODO: call gateway based on blockchain and call method of gateway based on schemaname?

		// TODO: gateway returns nft metadata

		// TODO: need to also verify owner in seperate gateway call (concurrent)?

		var metadata *gateways.NFTMetadata
		var err error
		switch schemaName {
		case "erc721":
			metadata, err = gateway.GetERC721NFTMetadata(contractAddress, tokenId)
		case "erc1155":
			metadata, err = gateway.GetERC1155NFTMetadata(contractAddress, tokenId)
		default:
			metadata = nil
			err = errors.New("invalida schema name")
		}

		if err != nil {
			logger.Err(err, "err on erc 1155")
			return nil, err
		}

		metaJson, _ := json.MarshalIndent(metadata, "", "  ")
		logger.Infof("result:\n%s\n", metaJson)

		nft := &entities.NFT{
			TokenId:         tokenId,
			Blockchain:      blockchain,
			ContractAddress: contractAddress,
			SchemaName:      schemaName,
			Owned:           false,
		}

		// gateway.VerifyERC1155Owner(contractAddress, address, tokenId)
		// gateway.VerifyERC721Owner(contractAddress, address, tokenId)

		return nft, nil
	}
}
