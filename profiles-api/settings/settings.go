package settings

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ISettings interface {
	CacheAddr() string
	CacheDB() int
	CachePassword() string
	CacheUseTLS() bool
	DatabaseURI() string
	Database() string
	StoreBlockchain() string
	StoreAddress() string
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
	SpotlightProfileAddress() string
}

type settings struct {
	redisAddr                string
	redisDB                  string
	redisPassword            string
	redisUseTLS              string
	mongoURI                 string
	mongoDB                  string
	storeBlockchain          string
	storeAddress             string
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
	spotlightProfileAddress  string
}

func NewSettings() ISettings {
	return &settings{
		redisAddr:                os.Getenv("REDIS_ADDR"),
		redisPassword:            os.Getenv("REDIS_PASSWORD"),
		redisDB:                  os.Getenv("REDIS_DB"),
		redisUseTLS:              os.Getenv("REDIS_USE_TLS"),
		mongoURI:                 os.Getenv("MONGO_URI"),
		mongoDB:                  os.Getenv("MONGO_DB"),
		storeBlockchain:          os.Getenv("STORE_BLOCKCHAIN"),
		storeAddress:             os.Getenv("STORE_ADDRESS"),
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
		spotlightProfileAddress:  os.Getenv("SPOTLIGHT_PROFILE_ADDRESS"),
	}
}

func (s *settings) CacheAddr() string {
	return s.redisAddr
}

func (s *settings) CacheDB() int {
	redisDB, err := strconv.Atoi(s.redisDB)

	if err != nil {
		panic(fmt.Errorf("invalid redis db value %v %w", s.redisDB, err))
	}

	return redisDB
}

func (s *settings) CachePassword() string {
	return s.redisPassword
}

func (s *settings) CacheUseTLS() bool {
	useTLS, err := strconv.ParseBool(s.redisUseTLS)

	if err != nil {
		panic(fmt.Errorf("invalid redis use tls value %v %w", s.redisUseTLS, err))
	}

	return useTLS
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

func (s *settings) SpotlightProfileAddress() string {
	return s.spotlightProfileAddress
}
