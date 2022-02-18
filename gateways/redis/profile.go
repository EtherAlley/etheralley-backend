package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/etheralley/etheralley-core-api/entities"
)

const ProfileNamespace = "profile"

func (g *Gateway) GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error) {
	profileString, err := g.client.Get(ctx, getFullKey(ProfileNamespace, address)).Result()

	if err != nil {
		return nil, err
	}

	profJson := &profileJson{}
	err = json.Unmarshal([]byte(profileString), profJson)

	if err != nil {
		return nil, err
	}

	profile := fromProfileJson(profJson)

	return profile, nil
}

func (g *Gateway) SaveProfile(ctx context.Context, profile *entities.Profile) error {
	profJson := toProfileJson(profile)

	bytes, err := json.Marshal(profJson)

	if err != nil {
		return err
	}

	_, err = g.client.Set(ctx, getFullKey(ProfileNamespace, profile.Address), bytes, time.Hour*24).Result()

	return err
}

func fromProfileJson(profileJson *profileJson) *entities.Profile {
	nfts := []entities.NonFungibleToken{}
	for _, nft := range *profileJson.NonFungibleTokens {
		nfts = append(nfts, entities.NonFungibleToken{
			TokenId: nft.TokenId,
			Contract: &entities.Contract{
				Blockchain: nft.Contract.Blockchain,
				Address:    nft.Contract.Address,
				Interface:  nft.Contract.Interface,
			},
			Balance: nft.Balance,
			Metadata: &entities.NonFungibleMetadata{
				Name:        nft.Metadata.Name,
				Description: nft.Metadata.Description,
				Image:       nft.Metadata.Image,
				Attributes:  nft.Metadata.Attributes,
			},
		})
	}
	tokens := []entities.FungibleToken{}
	for _, token := range *profileJson.FungibleTokens {
		tokens = append(tokens, entities.FungibleToken{
			Contract: &entities.Contract{
				Blockchain: token.Contract.Blockchain,
				Address:    token.Contract.Address,
				Interface:  token.Contract.Interface,
			},
			Balance: token.Balance,
			Metadata: &entities.FungibleMetadata{
				Name:     token.Metadata.Name,
				Symbol:   token.Metadata.Symbol,
				Decimals: token.Metadata.Decimals,
			},
		})
	}
	stats := []entities.Statistic{}
	for _, stat := range *profileJson.Statistics {
		stats = append(stats, entities.Statistic{
			Type: stat.Type,
			Contract: &entities.Contract{
				Blockchain: stat.Contract.Blockchain,
				Address:    stat.Contract.Address,
				Interface:  stat.Contract.Interface,
			},
			Data: &stat.Data,
		})
	}
	interactions := []entities.Interaction{}
	for _, interaction := range *profileJson.Interactions {
		interactions = append(interactions, entities.Interaction{
			Type: interaction.Type,
			Transaction: &entities.Transaction{
				Blockchain: interaction.Transaction.Blockchain,
				Id:         interaction.Transaction.Id,
			},
			Timestamp: interaction.Timestamp,
		})
	}
	return &entities.Profile{
		Address:           profileJson.Address,
		ENSName:           profileJson.ENSName,
		NonFungibleTokens: &nfts,
		FungibleTokens:    &tokens,
		Statistics:        &stats,
		Interactions:      &interactions,
	}
}

func toProfileJson(profile *entities.Profile) *profileJson {
	nfts := []nonFungibleTokenJson{}
	for _, nft := range *profile.NonFungibleTokens {
		nfts = append(nfts, nonFungibleTokenJson{
			TokenId: nft.TokenId,
			Contract: &contractJson{
				Blockchain: nft.Contract.Blockchain,
				Address:    nft.Contract.Address,
				Interface:  nft.Contract.Interface,
			},
			Balance: nft.Balance,
			Metadata: &nonFungibleMetadataJson{
				Name:        nft.Metadata.Name,
				Description: nft.Metadata.Description,
				Image:       nft.Metadata.Image,
				Attributes:  nft.Metadata.Attributes,
			},
		})
	}
	tokens := []fungibleTokenJson{}
	for _, token := range *profile.FungibleTokens {
		tokens = append(tokens, fungibleTokenJson{
			Contract: &contractJson{
				Blockchain: token.Contract.Blockchain,
				Address:    token.Contract.Address,
				Interface:  token.Contract.Interface,
			},
			Balance: token.Balance,
			Metadata: &fungibleMetadataJson{
				Name:     token.Metadata.Name,
				Symbol:   token.Metadata.Symbol,
				Decimals: token.Metadata.Decimals,
			},
		})
	}
	stats := []statisticJson{}
	for _, stat := range *profile.Statistics {
		stats = append(stats, statisticJson{
			Type: stat.Type,
			Contract: &contractJson{
				Blockchain: stat.Contract.Blockchain,
				Address:    stat.Contract.Address,
				Interface:  stat.Contract.Interface,
			},
			Data: stat.Data,
		})
	}
	interactions := []interactionJson{}
	for _, interaction := range *profile.Interactions {
		interactions = append(interactions, interactionJson{
			Type: interaction.Type,
			Transaction: &transactionJson{
				Blockchain: interaction.Transaction.Blockchain,
				Id:         interaction.Transaction.Id,
			},
			Timestamp: interaction.Timestamp,
		})
	}
	return &profileJson{
		Address:           profile.Address,
		ENSName:           profile.ENSName,
		NonFungibleTokens: &nfts,
		FungibleTokens:    &tokens,
		Statistics:        &stats,
		Interactions:      &interactions,
	}
}
