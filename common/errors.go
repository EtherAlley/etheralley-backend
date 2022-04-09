package common

import (
	"errors"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrValidation = errors.New("validation")
	ErrRetryable  = errors.New("retryable")
)
