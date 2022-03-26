package ethereum

import (
	"context"
	"errors"
	"math/big"

	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (gw *gateway) GetStoreListingInfo(ctx context.Context, ids *[]string) (*[]entities.ListingInfo, error) {
	client, err := gw.getClient(ctx, gw.settings.StoreBlockchain())

	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(gw.settings.StoreAddress())

	instance, err := contracts.NewEtherAlleyStore(contractAddress, client)

	if err != nil {
		return nil, err
	}

	idsArr := []*big.Int{}

	for _, id := range *ids {
		n := new(big.Int)
		n, ok := n.SetString(id, 10)

		if !ok {
			return nil, errors.New("err parsing id into big int")
		}

		idsArr = append(idsArr, n)
	}

	listingsArr, err := instance.GetListingBatch(&bind.CallOpts{}, idsArr)

	if err != nil {
		return nil, err
	}

	listings := []entities.ListingInfo{}
	for _, listing := range listingsArr {
		listings = append(listings, entities.ListingInfo{
			Purchasable:  listing.Purchasable,
			Transferable: listing.Transferable,
			Price:        listing.Price.String(),
			SupplyLimit:  listing.SupplyLimit.String(),
			BalanceLimit: listing.BalanceLimit.String(),
		})
	}

	return &listings, err
}

func (gw *gateway) GetStoreBalanceBatch(ctx context.Context, address string, ids *[]string) ([]*big.Int, error) {
	client, err := gw.getClient(ctx, gw.settings.StoreBlockchain())

	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(gw.settings.StoreAddress())

	instance, err := contracts.NewEtherAlleyStore(contractAddress, client)

	if err != nil {
		return nil, err
	}

	idsArr := []*big.Int{}
	for _, id := range *ids {
		n := new(big.Int)
		n, ok := n.SetString(id, 10)

		if !ok {
			return nil, errors.New("err parsing id into big int")
		}

		idsArr = append(idsArr, n)
	}

	accountsArr := make([]common.Address, len(*ids))
	for i := 0; i < len(*ids); i++ {
		accountsArr[i] = common.HexToAddress(address)
	}

	return instance.BalanceOfBatch(&bind.CallOpts{}, accountsArr, idsArr)

	// if err != nil {
	// 	return nil, err
	// }

	// balences := make([]string, len(*ids))
	// for i := 0; i < len(*ids); i++ {
	// 	balences[i] = balancesArr[i].String()
	// }

	// return &balences, err
}
