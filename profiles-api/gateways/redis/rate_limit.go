package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis_rate/v9"
)

func (gw *gateway) VerifyRateLimit(ctx context.Context, ipAddress string) error {
	res, err := gw.limiter.Allow(ctx, ipAddress, redis_rate.PerSecond(10))

	if err != nil {
		return fmt.Errorf("rate limit error %w", err)
	}

	if res.Remaining == 0 {
		return fmt.Errorf("rate limit exceeded %v", res.Limit)
	}

	return nil
}
