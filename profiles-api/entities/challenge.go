package entities

import (
	"fmt"
	"math/rand"
	"time"
)

// how long a signature will ultimately be valid for
const CHALLENGE_TTL = time.Hour * 24

// the message that will be presented to the user in their wallet when signing
const CHALLENGE_MESSAGE string = `
Please sign this message to prove you are the owner of this address: %v

Signing this message does not cost anything.

Timestamp: %v
Nonce: %v
`

type Challenge struct {
	Address string
	Message string
	TTL     time.Duration
	Expires time.Time
}

func NewChallenge(address string) *Challenge {
	now := time.Now()
	ttl := CHALLENGE_TTL
	nonce := fmt.Sprintf("%10d", rand.Intn(10000000000))
	return &Challenge{
		Address: address,
		Message: fmt.Sprintf(CHALLENGE_MESSAGE, address, now.Format(time.RFC3339), nonce),
		TTL:     ttl,
		Expires: now.Add(ttl),
	}
}
