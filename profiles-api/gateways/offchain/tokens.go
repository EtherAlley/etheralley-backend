package offchain

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

type tokenMetadata map[string]struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
}

// json metadata is read into memory on app init
func (gw *gateway) initTokenMetadata() error {
	for _, blockchain := range []common.Blockchain{common.ARBITRUM, common.ETHEREUM, common.POLYGON, common.OPTIMISM} {
		file, err := ioutil.ReadFile(fmt.Sprintf("profiles-api/gateways/offchain/assets/tokens/%v.json", blockchain))

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

type alchemyGetTokenBalancesReqJson struct {
	JSONRPC string    `json:"jsonrpc"`
	Method  string    `json:"method"`
	Params  *[]string `json:"params"`
	Id      string    `json:"id"`
}

type alchemyGetTokenBalancesResponseJson struct {
	JSONRPC string `json:"jsonrpc"`
	Id      string `json:"id"`
	Result  *struct {
		Address       string `json:"address"`
		TokenBalances *[]struct {
			ContractAddress string  `json:"contractAddress"`
			TokenBalance    *string `json:"tokenBalance"`
			Error           *string `json:"error"`
		} `json:"tokenBalances"`
	} `json:"result"`
}

// See https://docs.alchemy.com/alchemy/enhanced-apis/token-api/alchemy_gettokenbalances
// We call this alchemy gettokenbalances api to detect erc20 contracts that this user has a balance for
func (gw *gateway) GetFungibleContracts(ctx context.Context, address string) (*[]entities.Contract, error) {
	body, err := json.Marshal(alchemyGetTokenBalancesReqJson{
		JSONRPC: "2.0",
		Method:  "alchemy_getTokenBalances",
		Params:  &[]string{address, "DEFAULT_TOKENS"},
		Id:      "42",
	})

	if err != nil {
		return nil, fmt.Errorf("encode get all token contracts req body %w", err)
	}

	resp, err := gw.httpClient.Do(ctx, "POST", gw.settings.AlchemyEthereumURI(), bytes.NewBuffer(body), &common.HttpOptions{})

	if err != nil {
		return nil, fmt.Errorf("get all token contracts req %w", err)
	}

	respJson := &alchemyGetTokenBalancesResponseJson{}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(respJson)

	if err != nil {
		return nil, fmt.Errorf("decode get all contracts response %w", err)
	}

	contracts := []entities.Contract{}
	for _, result := range *respJson.Result.TokenBalances {
		if result.Error != nil || result.TokenBalance == nil {
			continue
		}

		balance, parseErr := strconv.ParseInt(strings.Split(*result.TokenBalance, "x")[1], 16, 64)

		if parseErr != nil || balance <= 0 {
			continue
		}

		contracts = append(contracts, entities.Contract{
			Blockchain: common.ETHEREUM,
			Address:    result.ContractAddress,
			Interface:  common.ERC20,
		})
	}

	return &contracts, nil
}
