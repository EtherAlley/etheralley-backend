package common

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Settings struct {
	Env           string
	Port          string
	RedisPort     string
	RedisDB       int
	RedisPassword string
	MongoURI      string
	MongoDB       string
}

func NewSettings() *Settings {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading settings")
	}

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		panic("Error parsing redis db value")
	}

	return &Settings{
		Env:           os.Getenv("ENV"),
		Port:          os.Getenv("PORT"),
		MongoURI:      os.Getenv("MONGO_URI"),
		MongoDB:       os.Getenv("MONGO_DB"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       redisDB,
	}
}

func (settings *Settings) IsDev() bool {
	return settings.Env == "dev"
}
