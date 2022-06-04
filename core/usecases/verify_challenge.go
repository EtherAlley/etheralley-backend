package usecases

import (
	"context"
	"errors"
	"fmt"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/core/gateways"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func NewVerifyChallenge(cacheGateway gateways.ICacheGateway) IVerifyChallengeUseCase {
	return &verifyChallengeUseCase{
		cacheGateway,
	}
}

type verifyChallengeUseCase struct {
	cacheGateway gateways.ICacheGateway
}

type IVerifyChallengeUseCase interface {
	// verify if the provided signature was signed with the correct address and signed the correct challenge message
	Do(ctx context.Context, input *VerifyChallengeInput) error
}

type VerifyChallengeInput struct {
	Address string `validate:"required,eth_addr"`
	SigHex  string `validate:"required"`
}

// get the current challenge message for the provided address
// hash the challenge message and use that to get the public key out of the signature
// compare the address from the signature with the provided address
// https://gist.github.com/dcb9/385631846097e1f59e3cba3b1d42f3ed#file-eth_sign_verify-go
func (uc *verifyChallengeUseCase) Do(ctx context.Context, input *VerifyChallengeInput) error {
	if err := cmn.ValidateStruct(input); err != nil {
		return err
	}

	if ok := common.IsHexAddress(input.Address); !ok {
		return errors.New("invalid address format")
	}

	challenge, err := uc.cacheGateway.GetChallengeByAddress(ctx, input.Address)

	if err != nil {
		return errors.New("no challenge for provided address")
	}

	msgBytes := []byte(challenge.Message)
	fromAddr := common.HexToAddress(input.Address)
	sig, err := hexutil.Decode(input.SigHex)

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
