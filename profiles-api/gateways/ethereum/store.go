package ethereum

import (
	"context"
	"fmt"
	"math/big"

	cmn "github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
	"github.com/etheralley/etheralley-backend/profiles-api/gateways/ethereum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (gw *gateway) GetStoreListingInfo(ctx context.Context, ids *[]string) (*[]entities.ListingInfo, error) {
	client, err := gw.getClient(ctx, gw.settings.StoreBlockchain())

	if err != nil {
		return nil, fmt.Errorf("listing info client %w", err)
	}

	contractAddress := common.HexToAddress(gw.settings.StoreAddress())

	instance, err := contracts.NewEtherAlleyStore(contractAddress, client)

	if err != nil {
		return nil, fmt.Errorf("listing info contract %w", err)
	}

	idsArr := []*big.Int{}

	for _, id := range *ids {
		n := new(big.Int)
		n, ok := n.SetString(id, 10)

		if !ok {
			return nil, fmt.Errorf("listing info parsing id %v", id)
		}

		idsArr = append(idsArr, n)
	}

	listingsArr, err := cmn.FunctionRetrier(ctx, func() ([]contracts.IEtherAlleyStoreTokenListing, error) {
		listingsArr, err := instance.GetListingBatch(&bind.CallOpts{}, idsArr)
		return listingsArr, gw.tryWrapRetryable(ctx, "listing info retry", err)
	})

	if err != nil {
		return nil, fmt.Errorf("listing info %w", err)
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

	return &listings, nil
}

func (gw *gateway) GetStoreBalanceBatch(ctx context.Context, address string, ids *[]string) ([]*big.Int, error) {
	client, err := gw.getClient(ctx, gw.settings.StoreBlockchain())

	if err != nil {
		return nil, fmt.Errorf("balance batch client %w", err)
	}

	contractAddress := common.HexToAddress(gw.settings.StoreAddress())

	instance, err := contracts.NewEtherAlleyStore(contractAddress, client)

	if err != nil {
		return nil, fmt.Errorf("balance batch contract %w", err)
	}

	idsArr := []*big.Int{}
	for _, id := range *ids {
		n := new(big.Int)
		n, ok := n.SetString(id, 10)

		if !ok {
			return nil, fmt.Errorf("balance batch parsing id %v", id)
		}

		idsArr = append(idsArr, n)
	}

	accountsArr := make([]common.Address, len(*ids))
	for i := 0; i < len(*ids); i++ {
		accountsArr[i] = common.HexToAddress(address)
	}

	return cmn.FunctionRetrier(ctx, func() ([]*big.Int, error) {
		balences, err := instance.BalanceOfBatch(&bind.CallOpts{}, accountsArr, idsArr)
		return balences, gw.tryWrapRetryable(ctx, "balance batch retry", err)
	})
}
