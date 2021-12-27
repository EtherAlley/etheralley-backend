package usecases

import (
	"context"
	"fmt"

	"github.com/eflem00/go-example-app/entities"
	"github.com/eflem00/go-example-app/gateways/redis"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type AuthenticationUseCase struct {
	challengeCache *redis.ChallengeCache
}

func NewAuthenticationUseCase(challengeCache *redis.ChallengeCache) *AuthenticationUseCase {
	return &AuthenticationUseCase{
		challengeCache,
	}
}

func (uc *AuthenticationUseCase) GetChallenge(ctx context.Context, address string) (*entities.Challenge, error) {
	challenge, err := uc.challengeCache.GetChallenge(ctx, address)

	if err != nil {
		challenge = entities.NewChallenge()
		err = uc.challengeCache.SetChallenge(ctx, address, challenge)
	}

	return challenge, err
}

// https://gist.github.com/dcb9/385631846097e1f59e3cba3b1d42f3ed#file-eth_sign_verify-go
func (uc *AuthenticationUseCase) VerifySignature(from, sigHex string, msgBytes []byte) bool {
	fromAddr := common.HexToAddress(from)

	sig, err := hexutil.Decode(sigHex)

	if err != nil || (sig[64] != 27 && sig[64] != 28) {
		return false
	}

	sig[64] -= 27

	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msgBytes), msgBytes)
	hash := crypto.Keccak256([]byte(msg))

	pubKey, err := crypto.SigToPub(hash, sig)

	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return fromAddr == recoveredAddr
}
