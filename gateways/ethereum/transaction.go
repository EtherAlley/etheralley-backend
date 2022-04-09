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
		return nil, err
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

		transaction, err := cmn.FunctionRetrier(ctx, func() (*types.Transaction, error) {
			tx, isPending, err := client.TransactionByHash(ctx, hash)

			if isPending {
				return nil, fmt.Errorf("get transaction transaction is pending %w", cmn.ErrRetryable)
			}

			return tx, tryWrapRetryable("get transaction", err)
		})

		if err != nil {
			txErr = err
			return
		}
		tx = transaction
	}()

	go func() {
		defer wg.Done()

		txRct, err := cmn.FunctionRetrier(ctx, func() (*types.Receipt, error) {
			txRct, err := client.TransactionReceipt(ctx, hash)
			return txRct, tryWrapRetryable("get transaction receipt", err)
		})

		if err != nil {
			headerErr = err
			return
		}

		hdr, err := client.HeaderByNumber(ctx, txRct.BlockNumber)

		headerErr = err
		header = hdr
	}()

	wg.Wait()

	if txErr != nil {
		return nil, txErr
	}

	if headerErr != nil {
		return nil, headerErr
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
		Timestamp: header.Time,
		From:      from.Hex(),
		To:        to,
		Data:      data,
		Value:     value,
	}

	return txData, nil
}
