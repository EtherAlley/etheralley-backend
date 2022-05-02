package offchain

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/etheralley/etheralley-core-api/entities"
)

type cryptoPunksMetadata map[string]struct {
	Name       string                   `json:"name"`
	Image      string                   `json:"image"`
	Attributes []map[string]interface{} `json:"attributes"`
}

const filename = "assets/cryptopunks/metadata.json"

// json metadata is read into memory on app init
func (gw *gateway) initPunkMetadata() {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(fmt.Errorf("could not read %v: %w", filename, err))
	}

	metadata := &cryptoPunksMetadata{}
	err = json.Unmarshal([]byte(file), metadata)

	if err != nil {
		panic(fmt.Errorf("could not unmarshal %v: %w", filename, err))
	}

	gw.cryptoPunkMetadata = metadata
}

func (gw *gateway) GetPunkMetadata(ctx context.Context, tokenId string) (*entities.NonFungibleMetadata, error) {
	punk, ok := (*gw.cryptoPunkMetadata)[tokenId]

	if !ok {
		return nil, fmt.Errorf("invalid crypto punk id")
	}

	return &entities.NonFungibleMetadata{
		Name:        punk.Name,
		Description: "",
		Image:       punk.Image,
		Attributes:  &punk.Attributes,
	}, nil
}
