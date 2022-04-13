package ethereum

import (
	"context"
	"fmt"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/ethereum/go-ethereum/common"
)

func (gw *gateway) GetAccountBalance(ctx context.Context, blockchain cmn.Blockchain, address string) (string, error) {
	client, err := gw.getClient(ctx, blockchain)

	if err != nil {
		return "", fmt.Errorf("account balance getClient %w", err)
	}

	addr := common.HexToAddress(address)

	balance, err := client.BalanceAt(ctx, addr, nil)

	if err != nil {
		return "", fmt.Errorf("account balance BalanceAt %w", err)
	}

	return balance.String(), nil
}
