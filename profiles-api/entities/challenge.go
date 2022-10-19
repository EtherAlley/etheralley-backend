package entities

import (
	"fmt"
	"time"
)

type Challenge struct {
	Address string
	Message string
}

func NewChallenge(address string) *Challenge {
	return &Challenge{
		Address: address,
		Message: fmt.Sprintf(`
		Hey there,
		
		This is a challenge message presented to you by etheralley.io.
		
		Please sign this message to prove you are the owner of this address
		
		Timestamp: %v
	`, time.Now().Format(time.RFC3339)),
	}
}
