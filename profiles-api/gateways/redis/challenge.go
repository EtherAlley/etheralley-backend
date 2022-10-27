package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

const ChallengeNamespace = "challenge"

func (g *gateway) GetChallengeByAddress(ctx context.Context, address string) (*entities.Challenge, error) {
	challengeStr, err := g.client.Get(ctx, getFullKey(ChallengeNamespace, address)).Result()

	if err != nil {
		return nil, fmt.Errorf("get challenge %w", err)
	}

	chalJson := &challengeJson{}
	err = json.Unmarshal([]byte(challengeStr), chalJson)

	if err != nil {
		return nil, fmt.Errorf("decode challenge %w", err)
	}

	challenge := fromChallengeJson(chalJson)

	return challenge, nil
}

func (g *gateway) SaveChallenge(ctx context.Context, challenge *entities.Challenge) error {
	challengeJson := toChallengeJson(challenge)
	bytes, err := json.Marshal(challengeJson)

	if err != nil {
		return fmt.Errorf("encode challenge %w", err)
	}

	_, err = g.client.Set(ctx, getFullKey(ChallengeNamespace, challenge.Address), bytes, challenge.TTL).Result()

	if err != nil {
		return fmt.Errorf("save challenge %w", err)
	}

	return nil
}
