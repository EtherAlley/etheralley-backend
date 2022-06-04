// always use layer 1 for ens resolution
// we also use the secondary uri for ens name resolution to avoid rate limiting
// we do this because failing to resolve the ens name is a very unpleseant experience for the end user
// they will see a 404 page implying their ens name does not exist

package ethereum

import (
	"context"
	"fmt"

	cmn "github.com/etheralley/etheralley-apis/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/v3"
)

func (gw *gateway) GetENSAddressFromName(ctx context.Context, name string) (string, error) {
	client, err := ethclient.DialContext(ctx, gw.settings.EthereumSecondaryURI())

	if err != nil {
		return "", fmt.Errorf("ens building client %w", err)
	}

	adr, err := cmn.FunctionRetrier(ctx, func() (common.Address, error) {
		adr, err := ens.Resolve(client, name)
		return adr, gw.tryWrapRetryable(ctx, "get address from ens name", err)
	})

	if err != nil {
		return "", fmt.Errorf("get address from ens name %w", err)
	}

	return adr.Hex(), nil
}

func (gw *gateway) GetENSNameFromAddress(ctx context.Context, address string) (string, error) {
	client, err := ethclient.DialContext(ctx, gw.settings.EthereumMainURI())

	if err != nil {
		return "", fmt.Errorf("ens building client %w", err)
	}

	return cmn.FunctionRetrier(ctx, func() (string, error) {
		name, err := ens.ReverseResolve(client, common.HexToAddress(address))
		return name, gw.tryWrapRetryable(ctx, "get ens name from address", err)
	})
}
