package entities

import "math/rand"

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

func (c *Challenge) Bytes() []byte {
	return []byte(c.Message)
}

func randString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, 32)
	for i := range bytes {
		bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(bytes)
}
