package redis

import (
	"context"
	"time"

	"github.com/etheralley/etheralley-core-api/entities"
)

const ChallengeNamespace = "challenge_"

func (g *Gateway) GetChallengeByAddress(ctx context.Context, address string) (*entities.Challenge, error) {
	msg, err := g.client.Get(ctx, GetFullKey(ChallengeNamespace, address)).Result()

	return &entities.Challenge{Address: address, Message: msg}, err
}

func (g *Gateway) SaveChallenge(ctx context.Context, challenge *entities.Challenge) error {
	_, err := g.client.Set(ctx, GetFullKey(ChallengeNamespace, challenge.Address), challenge.Message, time.Minute*5).Result()

	return err
}
