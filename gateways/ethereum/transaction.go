package ethereum

import (
	"context"
	"fmt"
	"sync"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (gw *gateway) GetTransactionData(ctx context.Context, transaction *entities.Transaction) (*entities.TransactionData, error) {
	client, err := gw.getClient(ctx, transaction.Blockchain)

	if err != nil {
		return nil, fmt.Errorf("tx client %w", err)
	}

	hash := common.HexToHash(transaction.Id)

	var tx *types.Transaction
	var txErr error
	var header *types.Header
	var headerErr error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		tx, txErr = cmn.FunctionRetrier(ctx, func() (*types.Transaction, error) {
			tx, isPending, err := client.TransactionByHash(ctx, hash)

			if isPending {
				return nil, fmt.Errorf("tx is pending %w", cmn.ErrRetryable)
			}

			return tx, gw.tryWrapRetryable(ctx, "tx retry", err)
		})
	}()

	go func() {
		defer wg.Done()

		txRct, err := cmn.FunctionRetrier(ctx, func() (*types.Receipt, error) {
			txRct, err := client.TransactionReceipt(ctx, hash)
			return txRct, gw.tryWrapRetryable(ctx, "tx receipt retry", err)
		})

		if err != nil {
			header, headerErr = nil, fmt.Errorf("tx receipt %w", err)
			return
		}

		header, headerErr = client.HeaderByNumber(ctx, txRct.BlockNumber)
	}()

	wg.Wait()

	if txErr != nil {
		return nil, fmt.Errorf("tx %w", err)
	}

	if headerErr != nil {
		return nil, fmt.Errorf("header %w", err)
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
		return nil, fmt.Errorf("sender %w", err)
	}

	value := tx.Value().String()

	txData := &entities.TransactionData{
		Timestamp: header.Time,
		From:      from.Hex(),
		To:        to,
		Data:      data,
		Value:     value,
	}

	return txData, nil
}
