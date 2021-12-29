package gateways

import (
	"context"

	"github.com/eflem00/go-example-app/entities"
)

type IDatabaseGateway interface {
	GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error)
	SaveProfile(ctx context.Context, profile *entities.Profile) error
}

type ICacheGateway interface {
	GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error)
	SaveProfile(ctx context.Context, profile *entities.Profile) error
	GetChallengeByAddress(ctx context.Context, address string) (*entities.Challenge, error)
	SaveChallenge(ctx context.Context, challenge *entities.Challenge) error
}
