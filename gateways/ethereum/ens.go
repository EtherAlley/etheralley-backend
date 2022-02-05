package ethereum

import (
	"bytes"
	"context"
	"errors"
	"strings"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/net/idna"
)

// same for all chains
var ENSContractAddress = common.HexToAddress("00000000000C2E074eC69A0dFb2997BA6C7d2e1e")

// https://eips.ethereum.org/EIPS/eip-137
// https://docs.ens.domains/dapp-developer-guide/resolving-names
func (gw *Gateway) GetENSAddressFromName(ctx context.Context, name string) (address string, err error) {
	client, err := gw.getClient(ctx, cmn.ETHEREUM) // awlays use layer 1 for ens resolution

	if err != nil {
		return
	}

	normalizedName, err := normalize(name)

	if err != nil {
		return
	}

	hash := nameHash(normalizedName)

	registry, err := contracts.NewEnsRegistry(ENSContractAddress, client)

	if err != nil {
		return
	}

	resolverAddress, err := registry.Resolver(nil, hash)

	if err != nil {
		return
	}

	resolver, err := contracts.NewEnsResolver(resolverAddress, client)

	if err != nil {
		return
	}

	adr, err := resolver.Addr(nil, hash)

	if err != nil {
		return
	}

	if bytes.Equal(adr.Bytes(), zeroAddress.Bytes()) {
		err = errors.New("resolved zero address")
		return
	}

	address = adr.Hex()

	return
}

// https://docs.ens.domains/contract-api-reference/name-processing#normalising-names
func normalize(name string) (string, error) {
	profile := idna.New(idna.MapForLookup(), idna.StrictDomainName(false), idna.Transitional(false))

	normalizedName, err := profile.ToUnicode(name)

	if err != nil {
		return normalizedName, err
	}

	if strings.HasPrefix(name, ".") && !strings.HasPrefix(normalizedName, ".") {
		return "." + normalizedName, nil
	}

	return normalizedName, nil
}

// https://docs.ens.domains/contract-api-reference/name-processing#hashing-names
func nameHash(name string) (hash [32]byte) {
	parts := strings.Split(name, ".")

	for i := len(parts) - 1; i >= 0; i-- {
		hash = nameHashPart(hash, parts[i])
	}

	return
}

func nameHashPart(currentHash [32]byte, name string) (hash [32]byte) {
	nameHash := crypto.Keccak256([]byte(name))

	newHash := crypto.Keccak256(append(currentHash[:], nameHash...))

	copy(hash[:], newHash[:32])

	return
}
