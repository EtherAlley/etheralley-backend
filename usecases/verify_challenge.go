package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/eflem00/go-example-app/gateways"
	"github.com/eflem00/go-example-app/gateways/redis"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type IVerifyChallengeUseCase interface {
	Go(ctx context.Context, address string, sigHex string) error
}

type VerifyChallengeUseCase struct {
	cacheGateway gateways.ICacheGateway
}

func NewVerifyChallengeUseCase(cacheGateway *redis.Gateway) *VerifyChallengeUseCase {
	return &VerifyChallengeUseCase{
		cacheGateway,
	}
}

// https://gist.github.com/dcb9/385631846097e1f59e3cba3b1d42f3ed#file-eth_sign_verify-go
func (uc *VerifyChallengeUseCase) Go(ctx context.Context, address string, sigHex string) error {
	challenge, err := uc.cacheGateway.GetChallengeByAddress(ctx, address)

	if err != nil {
		return errors.New("no challenge for provided address")
	}

	msgBytes := challenge.Bytes()
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
