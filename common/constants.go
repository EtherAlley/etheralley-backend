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

type Address = string

const (
	ZERO_ADDRESS Address = "0x0000000000000000000000000000000000000000"
)

// testnet erc20 addresses
const (
	UNI_GOERLI  Address = "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"
	LINK_GOERLI Address = "0x14b7ba66139c234b1be9a157d4f8b985b8a7f762"
	HEX_GOERLI  Address = "0x08249c12c66c76ea384cf851bb3e274a2bb1874a"
	DAI_GOERLI  Address = "0x11fe4b6ae13d2a6055c8d9cf65c55bac32b5d844"
	SHIB_GOERLI Address = "0xC5Ad32f66e0dd5FA75dB3FF839AB7783ce9f1a68"
	CRO_GOERLI  Address = "0x014EB3F44A9687450459a219f47C154158fc62aD"
)

// mainnet erc20 addresses
const (
	UNI_MAINNET  Address = "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"
	LINK_MAINNET Address = "0x514910771af9ca656af840dff83e8264ecf986ca"
	HEX_MAINNET  Address = "0x2b591e99afe9f32eaa6214f7b7629768c40eeb39"
	DAI_MAINNET  Address = "0x6b175474e89094c44da98b954eedeac495271d0f"
	SHIB_MAINNET Address = "0x95aD61b0a150d79219dCF64E1E6Cc01f0B64C4cE"
	CRO_MAINNET  Address = "0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b"
)
