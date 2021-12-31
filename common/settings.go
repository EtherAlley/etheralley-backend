package common

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Settings struct {
	Env             string
	Port            string
	RedisHost       string
	RedisPort       string
	RedisDB         int
	RedisPassword   string
	MongoUsername   string
	MongoPassword   string
	MongoHost       string
	MongoPort       string
	MongoAdminDB    string
	MongoDB         string
	InfuraChain     string
	InfuraProjectId string
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
		Env:             os.Getenv("ENV"),
		Port:            os.Getenv("PORT"),
		RedisHost:       os.Getenv("REDIS_HOST"),
		RedisPort:       os.Getenv("REDIS_PORT"),
		RedisPassword:   os.Getenv("REDIS_PASSWORD"),
		RedisDB:         redisDB,
		MongoUsername:   os.Getenv("MONGO_USERNAME"),
		MongoPassword:   os.Getenv("MONGO_PASSWORD"),
		MongoHost:       os.Getenv("MONGO_HOST"),
		MongoPort:       os.Getenv("MONGO_PORT"),
		MongoAdminDB:    os.Getenv("MONGO_ADMIN_DB"),
		MongoDB:         os.Getenv("MONGO_DB"),
		InfuraChain:     os.Getenv("INFURA_CHAIN"),
		InfuraProjectId: os.Getenv("INFURA_PROJECT_ID"),
	}
}

func (settings *Settings) IsDev() bool {
	return settings.Env == "dev"
}
