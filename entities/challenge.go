package entities

import "math/rand"

type Challenge struct {
	Message string `json:"message"`
}

func NewChallenge() *Challenge {
	return &Challenge{
		Message: randString(),
	}
}

func (c *Challenge) Bytes() []byte {
	return []byte(c.Message)
}

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func randString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	stringLen := 10
	bytes := make([]byte, stringLen)
	for i := range bytes {
		bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(bytes)
}
