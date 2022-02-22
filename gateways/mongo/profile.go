package mongo

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (g *Gateway) GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error) {
	profileBson := &profileBson{}

	err := g.profiles.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: address}}).Decode(profileBson)

	if err == mongo.ErrNoDocuments {
		return nil, common.ErrNotFound
	}

	profile := fromProfileBson(profileBson)

	return profile, err
}

func (g *Gateway) SaveProfile(ctx context.Context, profile *entities.Profile) error {
	profileBson := toProfileBson(profile)

	_, err := g.profiles.UpdateOne(ctx, bson.D{primitive.E{Key: "_id", Value: profile.Address}}, bson.D{primitive.E{Key: "$set", Value: profileBson}}, options.Update().SetUpsert(true))

	return err
}

func fromProfileBson(profileBson *profileBson) *entities.Profile {
	nfts := []entities.NonFungibleToken{}
	for _, nft := range *profileBson.NonFungibleTokens {
		nfts = append(nfts, entities.NonFungibleToken{
			TokenId: nft.TokenId,
			Contract: &entities.Contract{
				Blockchain: nft.Contract.Blockchain,
				Address:    nft.Contract.Address,
				Interface:  nft.Contract.Interface,
			},
		})
	}
	tokens := []entities.FungibleToken{}
	for _, token := range *profileBson.FungibleTokens {
		tokens = append(tokens, entities.FungibleToken{
			Contract: &entities.Contract{
				Blockchain: token.Contract.Blockchain,
				Address:    token.Contract.Address,
				Interface:  token.Contract.Interface,
			},
		})
	}
	stats := []entities.Statistic{}
	for _, stat := range *profileBson.Statistics {
		stats = append(stats, entities.Statistic{
			Type: stat.Type,
			Contract: &entities.Contract{
				Blockchain: stat.Contract.Blockchain,
				Address:    stat.Contract.Address,
				Interface:  stat.Contract.Interface,
			},
		})
	}
	interactions := []entities.Interaction{}
	for _, interaction := range *profileBson.Interactions {
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
		Address:           profileBson.Address,
		NonFungibleTokens: &nfts,
		FungibleTokens:    &tokens,
		Statistics:        &stats,
		Interactions:      &interactions,
	}
}

func toProfileBson(profile *entities.Profile) *profileBson {
	nfts := []nonFungibleTokenBson{}
	for _, nft := range *profile.NonFungibleTokens {
		nfts = append(nfts, nonFungibleTokenBson{
			TokenId: nft.TokenId,
			Contract: &contractBson{
				Blockchain: nft.Contract.Blockchain,
				Address:    nft.Contract.Address,
				Interface:  nft.Contract.Interface,
			},
		})
	}
	tokens := []fungibleTokenBson{}
	for _, token := range *profile.FungibleTokens {
		tokens = append(tokens, fungibleTokenBson{
			Contract: &contractBson{
				Blockchain: token.Contract.Blockchain,
				Address:    token.Contract.Address,
				Interface:  token.Contract.Interface,
			},
		})
	}
	stats := []statisticBson{}
	for _, stat := range *profile.Statistics {
		stats = append(stats, statisticBson{
			Type: stat.Type,
			Contract: &contractBson{
				Blockchain: stat.Contract.Blockchain,
				Address:    stat.Contract.Address,
				Interface:  stat.Contract.Interface,
			},
		})
	}
	interactions := []interactionBson{}
	for _, interaction := range *profile.Interactions {
		interactions = append(interactions, interactionBson{
			Type: interaction.Type,
			Transaction: &transactionBson{
				Blockchain: interaction.Transaction.Blockchain,
				Id:         interaction.Transaction.Id,
			},
			Timestamp: interaction.Timestamp,
		})
	}
	return &profileBson{
		Address:           profile.Address,
		NonFungibleTokens: &nfts,
		FungibleTokens:    &tokens,
		Statistics:        &stats,
		Interactions:      &interactions,
	}
}
