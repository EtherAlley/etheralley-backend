package http

import (
	"encoding/json"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/presenters"
)

type httpPresenter struct {
	logger common.ILogger
}

func NewPresenter(logger common.ILogger) presenters.IPresenter {
	return &httpPresenter{
		logger,
	}
}

func render(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func renderNoBody(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

func toProfileJson(profile *entities.Profile) *profileJson {
	return &profileJson{
		Address:           profile.Address,
		ENSName:           profile.ENSName,
		NonFungibleTokens: toNonFungibleTokensJson(profile.NonFungibleTokens),
		FungibleTokens:    toFungibleTokensJson(profile.FungibleTokens),
		Statistics:        toStatisticsJson(profile.Statistics),
		Interactions:      toInteractionsJson(profile.Interactions),
	}
}

func toNonFungibleTokensJson(nfts *[]entities.NonFungibleToken) *[]nonFungibleTokenJson {
	nftsJson := []nonFungibleTokenJson{}

	for _, nft := range *nfts {
		nftsJson = append(nftsJson, *toNonFungibleJson(&nft))
	}

	return &nftsJson
}

func toFungibleTokensJson(tokens *[]entities.FungibleToken) *[]fungibleTokenJson {
	tokensJson := []fungibleTokenJson{}

	for _, token := range *tokens {
		tokensJson = append(tokensJson, *toFungibleJson(&token))
	}

	return &tokensJson
}

func toStatisticsJson(stats *[]entities.Statistic) *[]statisticJson {
	statsJson := []statisticJson{}

	for _, stat := range *stats {
		statsJson = append(statsJson, *toStatisticJson(&stat))
	}

	return &statsJson
}

func toInteractionsJson(interactions *[]entities.Interaction) *[]interactionJson {
	interactionsJson := []interactionJson{}

	for _, interaction := range *interactions {
		interactionsJson = append(interactionsJson, *toInteractionJson(&interaction))
	}

	return &interactionsJson
}

func toFungibleJson(token *entities.FungibleToken) *fungibleTokenJson {
	return &fungibleTokenJson{
		Contract: toContractJson(token.Contract),
		Balance:  token.Balance,
		Metadata: &fungibleMetadataJson{
			Name:     token.Metadata.Name,
			Symbol:   token.Metadata.Symbol,
			Decimals: token.Metadata.Decimals,
		},
	}
}

func toNonFungibleJson(nft *entities.NonFungibleToken) *nonFungibleTokenJson {
	var metadata *nonFungibleMetadataJson
	if nft.Metadata != nil {
		metadata = &nonFungibleMetadataJson{
			Name:        nft.Metadata.Name,
			Description: nft.Metadata.Description,
			Image:       nft.Metadata.Image,
			Attributes:  nft.Metadata.Attributes,
		}
	}
	return &nonFungibleTokenJson{
		Contract: toContractJson(nft.Contract),
		TokenId:  nft.TokenId,
		Balance:  nft.Balance,
		Metadata: metadata,
	}
}

func toStatisticJson(stat *entities.Statistic) *statisticJson {
	return &statisticJson{
		Contract: toContractJson(stat.Contract),
		Type:     stat.Type,
		Data:     stat.Data,
	}
}

func toContractJson(contract *entities.Contract) *contractJson {
	return &contractJson{
		Blockchain: contract.Blockchain,
		Address:    contract.Address,
		Interface:  contract.Interface,
	}
}

func toInteractionJson(interaction *entities.Interaction) *interactionJson {
	return &interactionJson{
		Transaction: toTransactionJson(interaction.Transaction),
		Type:        interaction.Type,
		Timestamp:   interaction.Timestamp,
	}
}

func toTransactionJson(transaction *entities.Transaction) *transactiontJson {
	return &transactiontJson{
		Blockchain: transaction.Blockchain,
		Id:         transaction.Id,
	}
}

type profileJson struct {
	Address           string                  `json:"address"`
	ENSName           string                  `json:"ens_name"`
	NonFungibleTokens *[]nonFungibleTokenJson `json:"non_fungible_tokens"`
	FungibleTokens    *[]fungibleTokenJson    `json:"fungible_tokens"`
	Statistics        *[]statisticJson        `json:"statistics"`
	Interactions      *[]interactionJson      `json:"interactions"`
}

type contractJson struct {
	Blockchain common.Blockchain `json:"blockchain"`
	Address    string            `json:"address"`
	Interface  common.Interface  `json:"interface"`
}

type fungibleTokenJson struct {
	Contract *contractJson         `json:"contract"`
	Balance  string                `json:"balance"`
	Metadata *fungibleMetadataJson `json:"metadata"`
}

type fungibleMetadataJson struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
}

type nonFungibleTokenJson struct {
	Contract *contractJson            `json:"contract"`
	TokenId  string                   `json:"token_id"`
	Balance  string                   `json:"balance"`
	Metadata *nonFungibleMetadataJson `json:"metadata"`
}

type nonFungibleMetadataJson struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Image       string                    `json:"image"`
	Attributes  *[]map[string]interface{} `json:"attributes"`
}

type statisticJson struct {
	Type     common.StatisticType `json:"type"`
	Contract *contractJson        `json:"contract"`
	Data     interface{}          `json:"data"`
}

type interactionJson struct {
	Transaction *transactiontJson  `json:"transaction"`
	Type        common.Interaction `json:"type"`
	Timestamp   uint64             `json:"timestamp"`
}

type transactiontJson struct {
	Id         string            `json:"id"`
	Blockchain common.Blockchain `json:"blockchain"`
}
