package common

import (
	"os"

	"github.com/joho/godotenv"
)

type Settings struct {
	Env        string
	Port       string
	RedisPort  string
	MongoDBURI string
}

func NewSettings() *Settings {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading settings")
	}

	return &Settings{
		Env:        os.Getenv("ENV"),
		Port:       os.Getenv("PORT"),
		RedisPort:  os.Getenv("REDIS_PORT"),
		MongoDBURI: os.Getenv("MONGODB_URI"),
	}
}

func (settings *Settings) IsDev() bool {
	return settings.Env == "dev"
}
