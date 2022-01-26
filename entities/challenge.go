package entities

import (
	"fmt"
	"math/rand"
)

type Challenge struct {
	Address string `json:"-"`
	Message string `json:"message"`
}

func NewChallenge(address string) *Challenge {
	return &Challenge{
		Address: address,
		Message: randString(),
	}
}

func randString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, 32)
	for i := range bytes {
		bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return fmt.Sprintf(`
		Hey there,
		
		This is a challenge message presented to you by etheralley.io.
		
		Please sign this message to prove you are the owner of this address
		
		Included is a random string unique to this request: %v
	`, string(bytes))
}
