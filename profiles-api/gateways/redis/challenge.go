package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

const ChallengeNamespace = "challenge"

func (g *gateway) GetChallengeByAddress(ctx context.Context, address string) (*entities.Challenge, error) {
	msg, err := g.client.Get(ctx, getFullKey(ChallengeNamespace, address)).Result()

	if err != nil {
		return nil, fmt.Errorf("get challenge %w", err)
	}

	return &entities.Challenge{Address: address, Message: msg}, nil
}

func (g *gateway) SaveChallenge(ctx context.Context, challenge *entities.Challenge) error {
	_, err := g.client.Set(ctx, getFullKey(ChallengeNamespace, challenge.Address), challenge.Message, time.Minute*5).Result()

	if err != nil {
		return fmt.Errorf("save challenge %w", err)
	}

	return nil
}
