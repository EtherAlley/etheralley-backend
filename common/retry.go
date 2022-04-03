package common

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"
)

var MAX_RETRIES float64 = 4

// TODO:
// - Only retry on specific errors
//
// Invoke a function with exponential backoff retries when an error is encountered.
// FunctionRetrier acceps a function fn with any signature and an arbitrary number of args.
// FunctionRetrier will invoke fn with args and return the cast result T and an error
// This was originally implemented to solve us getting rate limited by the alchemy APIs
// See: https://docs.alchemy.com/alchemy/documentation/rate-limits#retries
func FunctionRetrier[T any](ctx context.Context, logger ILogger, fn any, args ...any) (result T, err error) {
	var tries float64 = 0
	for {
		select {
		case <-ctx.Done():
			return result, fmt.Errorf("exceeded context deadline: %w", err)
		default:
			if tries >= MAX_RETRIES {
				return result, fmt.Errorf("exceeded max retries: %w", err)
			}

			result, err = invoker[T](fn, args...)

			if err == nil {
				return result, nil
			}

			// sleep a power of two seconds + a random number of seconds between 0 and 1
			sleepSecs := (time.Second * time.Duration(math.Pow(2, tries))) + time.Duration((rand.Float64() * float64(time.Second)))

			logger.Debugf(ctx, "retrier failed on try %v with err %v sleeping %v", tries, sleepSecs, err)

			tries++

			time.Sleep(sleepSecs)
		}
	}
}

func invoker[T any](fn any, args ...any) (result T, err error) {
	fnValue := reflect.ValueOf(fn)
	arguments := []reflect.Value{}
	for _, arg := range args {
		arguments = append(arguments, reflect.ValueOf(arg))
	}

	fnResults := fnValue.Call(arguments)

	fnResult := fnResults[0].Interface()
	fnErr := fnResults[1].Interface()

	if fnErr != nil {
		err = fnErr.(error)
		return result, err
	}

	result = fnResult.(T)
	return result, err
}
