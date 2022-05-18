package common

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type ISettings interface {
	Appname() string
	Hostname() string
	Env() string
	IsDev() bool
	Port() string
	CacheAddr() string
	CacheDB() int
	CachePassword() string
	DatabaseURI() string
	Database() string
	StoreBlockchain() string
	StoreAddress() string
	StoreImageURI() string
	// This ethereum uri should be used for most workflows.
	EthereumMainURI() string
	// This ethereum secondary uri is intended to offload traffic from the main uri and avoid rate limiting for critical paths.
	// this URI should be used sparringly for critical workflows.
	EthereumSecondaryURI() string
	PolygonMainURI() string
	ArbitrumMainURI() string
	OptimismMainURI() string
	ENSMetadataURI() string
	// This ethereum uri is specific to alchemy and is intended for alchemy only apis (e.g. get all nfts).
	AlchemyEthereumURI() string
	CryptoKittiesMetadataURI() string
	IPFSURI() string
	TheGraphURI() string
	TheGraphHostedURI() string
	DefaultTokenAddresses() []string
}

type settings struct {
	appname                  string
	hostname                 string
	instanceID               string
	env                      string
	port                     string
	redisAddr                string
	redisDB                  string
	redisPassword            string
	mongoURI                 string
	mongoDB                  string
	storeBlockchain          string
	storeAddress             string
	storeImageURI            string
	ethereumMainURI          string
	ethereumSecondaryURI     string
	polygonMainURI           string
	arbitrumMainURI          string
	optimismMainURI          string
	alchemyEthereumURI       string
	ensMetadataURI           string
	cryptoKittiesMetadataURI string
	ipfsURI                  string
	theGraphURI              string
	theGraphHostedURI        string
	defaultTokenAddresses    string
}

func NewSettings() ISettings {
	env := os.Getenv("ENV")

	if env == "dev" {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading settings")
		}
	}

	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}

	return &settings{
		env:                      env,
		appname:                  "core-api",
		hostname:                 hostname,
		port:                     os.Getenv("PORT"),
		redisAddr:                os.Getenv("REDIS_ADDR"),
		redisPassword:            os.Getenv("REDIS_PASSWORD"),
		redisDB:                  os.Getenv("REDIS_DB"),
		mongoURI:                 os.Getenv("MONGO_URI"),
		mongoDB:                  os.Getenv("MONGO_DB"),
		storeBlockchain:          os.Getenv("STORE_BLOCKCHAIN"),
		storeAddress:             os.Getenv("STORE_ADDRESS"),
		storeImageURI:            os.Getenv("STORE_IMAGE_URI"),
		ethereumMainURI:          os.Getenv("ETHEREUM_MAIN_URI"),
		ethereumSecondaryURI:     os.Getenv("ETHEREUM_SECONDARY_URI"),
		polygonMainURI:           os.Getenv("POLYGON_MAIN_URI"),
		arbitrumMainURI:          os.Getenv("ARBITRUM_MAIN_URI"),
		optimismMainURI:          os.Getenv("OPTIMISM_MAIN_URI"),
		alchemyEthereumURI:       os.Getenv("ALCHEMY_ETHEREUM_URI"),
		ensMetadataURI:           os.Getenv("ENS_METADATA_URI"),
		cryptoKittiesMetadataURI: os.Getenv("CRYPTO_KITTIES_METADATA_URI"),
		ipfsURI:                  os.Getenv("IPFS_URI"),
		theGraphURI:              os.Getenv("THE_GRAPH_URI"),
		theGraphHostedURI:        os.Getenv("THE_GRAPH_HOSTED_URI"),
		defaultTokenAddresses:    os.Getenv("DEFAULT_TOKEN_ADDRESSES"),
	}
}
func (s *settings) Appname() string {
	return s.appname
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
		return 0
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

func (s *settings) StoreBlockchain() string {
	return s.storeBlockchain
}

func (s *settings) StoreAddress() string {
	return s.storeAddress
}

func (s *settings) StoreImageURI() string {
	return s.storeImageURI
}

func (s *settings) EthereumMainURI() string {
	return s.ethereumMainURI
}

func (s *settings) EthereumSecondaryURI() string {
	return s.ethereumSecondaryURI
}

func (s *settings) PolygonMainURI() string {
	return s.polygonMainURI
}

func (s *settings) ArbitrumMainURI() string {
	return s.arbitrumMainURI
}

func (s *settings) OptimismMainURI() string {
	return s.optimismMainURI
}

func (s *settings) AlchemyEthereumURI() string {
	return s.alchemyEthereumURI
}

func (s *settings) ENSMetadataURI() string {
	return s.ensMetadataURI
}

func (s *settings) CryptoKittiesMetadataURI() string {
	return s.cryptoKittiesMetadataURI
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

func (s *settings) DefaultTokenAddresses() []string {
	return strings.Split(s.defaultTokenAddresses, ",")
}
