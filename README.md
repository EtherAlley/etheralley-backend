# etheralley-core-api

[![Issues](https://img.shields.io/github/issues/EtherAlley/etheralley-web-interface.svg?style=for-the-badge)](https://github.com/EtherAlley/etheralley-web-interface/issues)
[![GPL-3.0 License](https://img.shields.io/github/license/EtherAlley/etheralley-web-interface.svg?style=for-the-badge)](https://github.com/EtherAlley/etheralley-web-interface/blob/main/LICENSE)

This reposiotry contains the core rest api for the EtherAlley platform

## Local development

### Prerequisites

These steps must be completed before proceeding

1. Download [Go](https://go.dev/dl/)
2. Download [Docker](https://www.docker.com/products/docker-desktop/)
3. Acquire API keys for communicating with various blockchains. e.g. [Alchemy](https://www.alchemy.com/)
4. Acquire an API key for communicating with subgraphs on [TheGraph](https://thegraph.com/en/)

### Setup

1. Start the Mongo database in a docker container
   ```sh
    make docker-mongo
   ```
2. Start the Redis cache in a docker container
   ```sh
    make docker-redis
   ```
3. Add a file named `.env` in the root of the project with the following environment variables. Anything with `<REPLACE>` must be replaced with the keys acquired in the prerequisite steps
   ```
    ENV=dev
    PORT=8080
    REDIS_ADDR=localhost:6379
    REDIS_DB=0
    REDIS_PASSWORD=
    MONGO_URI=mongodb://mongoadmin:secret@localhost:27017/
    MONGO_DB=etheralley
    ETHEREUM_URI=https://eth-goerli.alchemyapi.io/v2/<REPLACE>
    POLYGON_URI=https://polygon-mumbai.g.alchemy.com/v2/<REPLACE>
    ARBITRUM_URI=https://arb-rinkeby.g.alchemy.com/v2/<REPLACE>
    OPTIMISM_URI=https://opt-kovan.g.alchemy.com/v2/<REPLACE>
    THE_GRAPH_URI=https://gateway.thegraph.com/api/<REPLACE>/subgraphs/id
    THE_GRAPH_HOSTED_URI=https://api.thegraph.com/subgraphs/name
    STORE_BLOCKCHAIN=polygon
    STORE_ADDRESS=0x15EC5d87f2A810466aCbd761f38c35ae36523FE7
    STORE_IMAGE_URI=http://localhost:3000
    DEFAULT_TOKEN_ADDRESSES=0x1f9840a85d5af5bf1d1762f925bdaddc4201f984
    IPFS_URI=https://gateway.ipfs.io/ipfs/
    ENS_METADATA_URI=https://metadata.ens.domains/goerli
    CRYPTO_KITTIES_METADATA_URI=https://api.cryptokitties.co/kitties
   ```
4. Start the web service
   ```sh
    make run
   ```

## License

Distributed under the GNU General Public License v3.0. See [LICENSE](https://github.com/EtherAlley/etheralley-web-interface/blob/main/LICENSE) for more information.
