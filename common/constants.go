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
	CRYPTO_KITTIES      Interface = "CRYPTO_KITTIES"
	CRYPTO_PUNKS        Interface = "CRYPTO_PUNKS"
	SUSHISWAP_EXCHANGE  Interface = "SUSHISWAP_EXCHANGE"
	UNISWAP_V2_EXCHANGE Interface = "UNISWAP_V2_EXCHANGE"
	UNISWAP_V3_EXCHANGE Interface = "UNISWAP_V3_EXCHANGE"
	ROCKET_POOL         Interface = "ROCKET_POOL"
)

type Interaction = string

const (
	CONTRACT_CREATION Interaction = "CONTRACT_CREATION"
	SEND_ETHER        Interaction = "SEND_ETHER"
)

type StatisticType = string

const (
	SWAP  StatisticType = "SWAP"
	STAKE StatisticType = "STAKE"
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
	ZERO_ADDRESS               Address = "0x0000000000000000000000000000000000000000"
	ENS_BASE_REGISTRAR_ADDRESS Address = "0x57f1887a8bf19b14fc0df6fd9b2acc9af147ea85"
)

type TokenIds = string

const (
	STORE_PREMIUM     TokenIds = "1"
	STORE_BETA_TESTER TokenIds = "2"
)

const (
	// This number should never be lower than the total count returned by default profile, so that new users can always save their initial profile
	REGULAR_TOTAL_BADGE_COUNT uint = 26
	PREMIUM_TOTAL_BADGE_COUNT uint = 50
	// Reminder: making these numbers bigger has implications for the max badge limit above. When a new user tries to save their default profile for the first time they may fail due to being above the badge limit
	DEFAULT_TOKEN_CUTOFF uint = 6
	DEFAULT_NFT_CUTOFF   uint = 13
)
