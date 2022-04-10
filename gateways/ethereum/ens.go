package ethereum

import (
	"context"
	"fmt"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v3"
)

func (gw *gateway) GetENSAddressFromName(ctx context.Context, name string) (string, error) {
	client, err := gw.getClient(ctx, cmn.ETHEREUM) // awlays use layer 1 for ens resolution

	if err != nil {
		return "", fmt.Errorf("ens building client %w", err)
	}

	adr, err := cmn.FunctionRetrier(ctx, func() (common.Address, error) {
		adr, err := ens.Resolve(client, name)
		return adr, tryWrapRetryable("get address from ens name", err)
	})

	if err != nil {
		return "", fmt.Errorf("get address from ens name %w", err)
	}

	return adr.Hex(), nil
}

func (gw *gateway) GetENSNameFromAddress(ctx context.Context, address string) (string, error) {
	client, err := gw.getClient(ctx, cmn.ETHEREUM) // awlays use layer 1 for ens resolution

	if err != nil {
		return "", fmt.Errorf("ens building client %w", err)
	}

	return cmn.FunctionRetrier(ctx, func() (string, error) {
		name, err := ens.ReverseResolve(client, common.HexToAddress(address))
		return name, tryWrapRetryable("get ens name from address", err)
	})
}
