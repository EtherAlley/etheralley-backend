package common

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type ISettings interface {
	Hostname() string
	InstanceID() string
	Env() string
	IsDev() bool
	Port() string
	CacheAddr() string
	CacheDB() int
	CachePassword() string
	DatabaseURI() string
	Database() string
	EthereumURI() string
	PolygonURI() string
	ArbitrumURI() string
	OptimismURI() string
	ENSMetadataURI() string
	IPFSURI() string
	TheGraphURI() string
	TheGraphHostedURI() string
}

type settings struct {
	hostname          string
	instanceID        string
	env               string
	port              string
	redisAddr         string
	redisDB           string
	redisPassword     string
	mongoURI          string
	mongoDB           string
	ethereumURI       string
	polygonURI        string
	arbitrumURI       string
	optimismURI       string
	ensMetadataURI    string
	ipfsURI           string
	theGraphURI       string
	theGraphHostedURI string
}

func NewSettings() ISettings {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading settings")
	}

	// See https://github.com/go-chi/chi/blob/master/middleware/request_id.go#L46
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}

	var buf [12]byte
	var instanceID string
	for len(instanceID) < 10 {
		rand.Read(buf[:])
		instanceID = base64.StdEncoding.EncodeToString(buf[:])
		instanceID = strings.NewReplacer("+", "", "/", "").Replace(instanceID)
	}
	instanceID = instanceID[0:10]

	return &settings{
		hostname:          hostname,
		instanceID:        instanceID,
		env:               os.Getenv("ENV"),
		port:              os.Getenv("PORT"),
		redisAddr:         os.Getenv("REDIS_ADDR"),
		redisPassword:     os.Getenv("REDIS_PASSWORD"),
		redisDB:           os.Getenv("REDIS_DB"),
		mongoURI:          os.Getenv("MONGO_URI"),
		mongoDB:           os.Getenv("MONGO_DB"),
		ethereumURI:       os.Getenv("ETHEREUM_URI"),
		polygonURI:        os.Getenv("POLYGON_URI"),
		arbitrumURI:       os.Getenv("ARBITRUM_URI"),
		optimismURI:       os.Getenv("OPTIMISM_URI"),
		ensMetadataURI:    os.Getenv("ENS_METADATA_URI"),
		ipfsURI:           os.Getenv("IPFS_URI"),
		theGraphURI:       os.Getenv("THE_GRAPH_URI"),
		theGraphHostedURI: os.Getenv("THE_GRAPH_HOSTED_URI"),
	}
}

func (s *settings) Hostname() string {
	return s.hostname
}

func (s *settings) InstanceID() string {
	return s.instanceID
}

func (s *settings) Env() string {
	return s.env
}

func (s *settings) IsDev() bool {
	return s.env == "dev"
}

func (s *settings) Port() string {
	return s.port
}

func (s *settings) CacheAddr() string {
	return s.redisAddr
}

func (s *settings) CacheDB() int {
	redisDB, err := strconv.Atoi(s.redisDB)

	if err != nil {
		return 1
	}

	return redisDB
}

func (s *settings) CachePassword() string {
	return s.redisPassword
}

func (s *settings) DatabaseURI() string {
	return s.mongoURI
}

func (s *settings) Database() string {
	return s.mongoDB
}

func (s *settings) EthereumURI() string {
	return s.ethereumURI
}

func (s *settings) PolygonURI() string {
	return s.polygonURI
}

func (s *settings) ArbitrumURI() string {
	return s.arbitrumURI
}

func (s *settings) OptimismURI() string {
	return s.optimismURI
}

func (s *settings) ENSMetadataURI() string {
	return s.ensMetadataURI
}

func (s *settings) IPFSURI() string {
	return s.ipfsURI
}

func (s *settings) TheGraphURI() string {
	return s.theGraphURI
}

func (s *settings) TheGraphHostedURI() string {
	return s.theGraphHostedURI
}
