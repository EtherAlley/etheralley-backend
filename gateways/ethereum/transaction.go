package ethereum

import (
	"context"
	"errors"
	"sync"

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

	var tx *types.Transaction
	var txErr error
	var header *types.Header
	var headerErr error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		transaction, isPending, err := client.TransactionByHash(ctx, hash)

		if err != nil {
			txErr = err
			return
		}

		if isPending {
			txErr = errors.New("pending transaction")
			return
		}
		tx = transaction
	}()

	go func() {
		defer wg.Done()

		txRct, err := client.TransactionReceipt(ctx, hash)

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
