package thegraph

import (
	"context"
	"fmt"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

type stakeJson struct {
	TotalRewards string `json:"total_rewards"`
}

func (gw *gateway) GetStake(ctx context.Context, address string, contract *entities.Contract) (interface{}, error) {
	uri := gw.settings.TheGraphURI()

	switch contract.Blockchain {
	case common.ETHEREUM:
		switch contract.Interface {
		case common.ROCKET_POOL:
			return gw.getRocketPoolStake(ctx, fmt.Sprintf("%v/S9ihna8D733WTEShJ1KctSTCvY1VJ7gdVwhUujq4Ejo", uri), address)
		}
	}

	return nil, fmt.Errorf("invalid subgraph for blockchain %v and interface %v", contract.Blockchain, contract.Interface)
}

type rocketPoolStakeQuery = struct {
	Stakers []struct {
		TotalEthRewards string `graphql:"totalETHRewards" json:"totalETHRewards"`
	} `graphql:"stakers(where: { id: $id })" json:"stakers"`
}

func (gw *gateway) getRocketPoolStake(ctx context.Context, url string, address string) (*stakeJson, error) {
	query := &rocketPoolStakeQuery{}
	variables := map[string]interface{}{
		"id": address,
	}
	err := gw.graphClient.Query(ctx, url, query, variables)

	if err != nil {
		return nil, fmt.Errorf("getRocketPoolStake query %w", err)
	}

	if len(query.Stakers) == 0 {
		return &stakeJson{
			TotalRewards: "0",
		}, nil
	}

	stake := query.Stakers[len(query.Stakers)-1]

	return &stakeJson{
		TotalRewards: stake.TotalEthRewards,
	}, nil
}
