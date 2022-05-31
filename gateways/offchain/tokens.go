package offchain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

type tokenMetadata map[string]struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
}

// json metadata is read into memory on app init
func (gw *gateway) initTokenMetadata() error {
	for _, blockchain := range []common.Blockchain{common.ARBITRUM, common.ETHEREUM, common.POLYGON, common.OPTIMISM} {
		file, err := ioutil.ReadFile(fmt.Sprintf("assets/tokens/%v.json", blockchain))

		if err != nil {
			return fmt.Errorf("could not read %v: %w", blockchain, err)
		}

		metadata := &tokenMetadata{}
		err = json.Unmarshal([]byte(file), metadata)

		if err != nil {
			return fmt.Errorf("could not unmarshal %v: %w", blockchain, err)
		}

		(*gw.tokenMetadata)[blockchain] = *metadata
	}

	return nil
}

func (gw *gateway) GetFungibleMetadata(ctx context.Context, contract *entities.Contract) (*entities.FungibleMetadata, error) {
	metadata, ok := (*gw.tokenMetadata)[contract.Blockchain]

	if !ok {
		return nil, errors.New("unsupported blockchain")
	}

	tokenMetadata, ok := metadata[strings.ToLower(contract.Address)]

	if !ok {
		return nil, errors.New("could not find token")
	}

	return &entities.FungibleMetadata{
		Name:     &tokenMetadata.Name,
		Decimals: &tokenMetadata.Decimals,
		Symbol:   &tokenMetadata.Symbol,
	}, nil
}
