package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func NewVerifyChallengeUseCase(cacheGateway *redis.Gateway) VerifyChallengeUseCase {
	return VerifyChallenge(cacheGateway)
}

// get the challenge message for the provided address
// hash the challenge message and get the public key out of the signature
// compare the address from the signature with the provided address
// https://gist.github.com/dcb9/385631846097e1f59e3cba3b1d42f3ed#file-eth_sign_verify-go
func VerifyChallenge(cacheGateway gateways.ICacheGateway) VerifyChallengeUseCase {
	return func(ctx context.Context, address string, sigHex string) error {
		challenge, err := cacheGateway.GetChallengeByAddress(ctx, address)

		if err != nil {
			return errors.New("no challenge for provided address")
		}

		msgBytes := []byte(challenge.Message)
		fromAddr := common.HexToAddress(address)
		sig, err := hexutil.Decode(sigHex)

		if err != nil || (sig[64] != 27 && sig[64] != 28) {
			return errors.New("invalid signature format")
		}

		sig[64] -= 27
		msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msgBytes), msgBytes)
		hash := crypto.Keccak256([]byte(msg))
		pubKey, err := crypto.SigToPub(hash, sig)

		if err != nil {
			return errors.New("error getting public key from signature")
		}

		recoveredAddr := crypto.PubkeyToAddress(*pubKey)

		if fromAddr != recoveredAddr {
			return errors.New("address mismatch")
		}

		return nil
	}
}
