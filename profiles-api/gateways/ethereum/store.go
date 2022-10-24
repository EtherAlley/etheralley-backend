package ethereum

import (
	"context"
	"fmt"
	"math/big"

	cmn "github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/gateways/ethereum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

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
