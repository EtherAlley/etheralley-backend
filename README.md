# etheralley-backend

The backend for the EtherAlley platform

[![Profiles API Build](https://github.com/etheralley/etheralley-backend/actions/workflows/build-profiles-api.yml/badge.svg)](https://github.com/etheralley/etheralley-backend/actions/workflows/build-profiles-api.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/etheralley/etheralley-backend/common)](https://goreportcard.com/report/github.com/etheralley/etheralley-backend/common)
[![Go Report Card](https://goreportcard.com/badge/github.com/etheralley/etheralley-backend/profiles-api)](https://goreportcard.com/report/github.com/etheralley/etheralley-backend/profiles-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/etheralley/etheralley-backend/daos-api)](https://goreportcard.com/report/github.com/etheralley/etheralley-backend/daos-api)
[![GPL-3.0 License](https://img.shields.io/github/license/EtherAlley/etheralley-backend.svg)](https://github.com/etheralley/etheralley-backend/profiles-api-api/blob/main/LICENSE)

## Repository Structure

This repository leverages the new [go workspace](https://go.dev/doc/tutorial/workspaces) feature to operate as a monorepo. There is currently two executable modules and a single library module.

- `profiles-api` is a rest api encapsulating the profiles domain
- `daos-api` is a rest api encapsulating the daos domain
- `common` is a library for command and shareable code

## Local Setup

1. Download [Go v18.x](https://go.dev/dl/)

2. Download [Docker](https://www.docker.com/products/docker-desktop/)

3. Acquire API keys for communicating with various blockchains. e.g. [Alchemy](https://www.alchemy.com/)

4. Acquire an API key for communicating with subgraphs on [TheGraph](https://thegraph.com/en/)

5. Start the Mongo database in a docker container
   ```sh
    make docker-run-mongo
   ```
6. Start the Redis cache in a docker container
   ```sh
    make docker-run-redis
   ```
7. Add a folder named `.env` in the root of the project with a file named `.env.profiles-api.dev` in the root of the project with the following environment variables. Anything with `<REPLACE>` must be replaced with the keys acquired in the previouse steps. `ETHEREUM_MAIN_URI`, `ETHEREUM_SECONDARY_URI` and `ALCHEMY_ETHEREUM_URI` can all point to the same uri on local and things should still behave normally.

   ```
   ENV=dev
   PORT=8080
   REDIS_ADDR=localhost:6379
   REDIS_DB=0
   REDIS_PASSWORD=
   REDIS_USE_TLS=false
   MONGO_URI=mongodb://mongoadmin:secret@localhost:27017/
   MONGO_DB=etheralley
   ETHEREUM_MAIN_URI=https://eth-goerli.alchemyapi.io/v2/<REPLACE>
   POLYGON_MAIN_URI=https://polygon-mumbai.g.alchemy.com/v2/<REPLACE>
   ARBITRUM_MAIN_URI=https://arb-rinkeby.g.alchemy.com/v2/<REPLACE>
   OPTIMISM_MAIN_URI=https://opt-kovan.g.alchemy.com/v2/<REPLACE>
   ETHEREUM_SECONDARY_URI=https://eth-goerli.alchemyapi.io/v2/<REPLACE>
   ALCHEMY_ETHEREUM_URI=https://eth-goerli.alchemyapi.io/v2/<REPLACE>
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

8. Start the profiles-api web service
   ```sh
    make run-profiles-api
   ```

## License

Distributed under the GNU General Public License v3.0. See [LICENSE](https://github.com/EtherAlley/etheralley-web-interface/blob/main/LICENSE) for more information.
