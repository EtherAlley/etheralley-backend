package usecases

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/eflem00/go-example-app/gateways/redis"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type AuthenticationUseCase struct {
	authCache *redis.AuthCache
}

func NewAuthenticationUseCase(authCache *redis.AuthCache) *AuthenticationUseCase {
	return &AuthenticationUseCase{
		authCache,
	}
}

func (uc *AuthenticationUseCase) GetChallengeMessage(ctx context.Context, address string) (string, error) {
	msg, err := uc.authCache.GetChallengeMessage(ctx, address)

	if err != nil {
		msg = randString()
		err = uc.authCache.SetChallengeMessage(ctx, address, msg)
		return msg, err
	}

	return msg, nil
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

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func randString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	stringLen := 10
	bytes := make([]byte, stringLen)
	for i := range bytes {
		bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(bytes)
}
