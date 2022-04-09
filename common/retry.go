package common

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const MAX_RETRIES float64 = 4

// Invoke a function with exponential backoff when a retryable error is encountered.
// If fn returns an err that is wrapped in common.ErrRetyable, fn will be retried with exponential backoff.
// This was originally implemented to solve getting rate limited by the alchemy APIs
//
// See: https://docs.alchemy.com/alchemy/documentation/rate-limits#retries
func FunctionRetrier[T any](ctx context.Context, fn func() (T, error)) (result T, err error) {
	var tries float64 = 0
	for {
		select {
		case <-ctx.Done():
			return result, fmt.Errorf("exceeded context deadline: %w", err)
		default:
			if tries >= MAX_RETRIES {
				return result, fmt.Errorf("exceeded max retries: %w", err)
			}

			result, err = fn()

			if err == nil || !errors.Is(err, ErrRetryable) {
				return result, err
			}

			// sleep a power of two seconds + a random number of seconds between 0 and 1
			sleepSecs := (time.Second * time.Duration(math.Pow(2, tries))) + time.Duration((rand.Float64() * float64(time.Second)))

			// sleep a random interval to get sufficient staggering between all the concurrent calls we have to do
			// sleepSecs := time.Duration((rand.Float64() * float64(time.Second))) * 5

			tries++

			time.Sleep(sleepSecs)
		}
	}
}
