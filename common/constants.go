package common

type Blockchain = string

const (
	ETHEREUM Blockchain = "ethereum"
	POLYGON  Blockchain = "polygon"
	ARBITRUM Blockchain = "arbitrum"
	OPTIMISM Blockchain = "optimism"
)

type Interface = string

const (
	ERC721              Interface = "ERC721"
	ERC1155             Interface = "ERC1155"
	ERC20               Interface = "ERC20"
	ENS_REGISTRAR       Interface = "ENS_REGISTRAR"
	SUSHISWAP_EXCHANGE  Interface = "SUSHISWAP_EXCHANGE"
	UNISWAP_V2_EXCHANGE Interface = "UNISWAP_V2_EXCHANGE"
	UNISWAP_V3_EXCHANGE Interface = "UNISWAP_V3_EXCHANGE"
)

type Interaction = string

const (
	CONTRACT_CREATION Interaction = "CONTRACT_CREATION"
	SEND_ETHER        Interaction = "SEND_ETHER"
)

type StatisticType = string

const (
	SWAP StatisticType = "SWAP"
)

type BadgeType = string

const (
	NON_FUNGIBLE_TOKEN BadgeType = "non_fungible_tokens"
	FUNGIBLE_TOKEN     BadgeType = "fungible_tokens"
	STATISTICS         BadgeType = "statistics"
)

type AchievementType = string

const (
	INTERACTIONS AchievementType = "interactions"
)

type Address = string

const (
	ZERO_ADDRESS Address = "0x0000000000000000000000000000000000000000"
)

type TokenIds = string

const (
	STORE_PREMIUM     TokenIds = "1"
	STORE_BETA_TESTER TokenIds = "2"
)
