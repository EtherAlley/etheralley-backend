package ethereum

import (
	"context"
	"errors"

	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (gw *Gateway) GetTransactionData(ctx context.Context, transaction *entities.Transaction) (*entities.TransactionData, error) {
	client, err := gw.getClient(ctx, transaction.Blockchain)

	if err != nil {
		return nil, err
	}

	hash := common.HexToHash(transaction.Id)

	tx, isPending, err := client.TransactionByHash(ctx, hash)

	if err != nil {
		return nil, err
	}

	if isPending {
		return nil, errors.New("pending transaction")
	}

	data := tx.Data()

	var to *string
	if tx.To() != nil {
		address := tx.To()
		toHex := address.Hex()
		to = &toHex
	}

	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)

	if err != nil {
		return nil, err
	}

	value := tx.Value().String()

	txData := &entities.TransactionData{
		From:  from.Hex(),
		To:    to,
		Data:  data,
		Value: value,
	}

	return txData, nil
}
