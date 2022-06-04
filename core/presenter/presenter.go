package presenter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/etheralley/etheralley-apis/common"
	"github.com/etheralley/etheralley-apis/core/entities"
	"github.com/etheralley/etheralley-apis/core/settings"
)

func NewHttpPresenter(logger common.ILogger, settings settings.ISettings) IHttpPresenter {
	return &presenter{
		logger,
		settings,
	}
}

type IHttpPresenter interface {
	PresentBadRequest(http.ResponseWriter, *http.Request, error)
	PresentUnathorized(http.ResponseWriter, *http.Request, error)
	PresentNotFound(http.ResponseWriter, *http.Request, error)
	PresentTooManyRequests(http.ResponseWriter, *http.Request, error)
	PresentForbiddenRequest(http.ResponseWriter, *http.Request, error)
	PresentHealth(http.ResponseWriter, *http.Request)
	PresentChallenge(http.ResponseWriter, *http.Request, *entities.Challenge)
	PresentFungibleToken(http.ResponseWriter, *http.Request, *entities.FungibleToken)
	PresentNonFungibleToken(http.ResponseWriter, *http.Request, *entities.NonFungibleToken)
	PresentStatistic(http.ResponseWriter, *http.Request, *entities.Statistic)
	PresentInteraction(http.ResponseWriter, *http.Request, *entities.Interaction)
	PresentProfile(http.ResponseWriter, *http.Request, *entities.Profile)
	PresentSavedProfile(http.ResponseWriter, *http.Request)
	PresentTopProfiles(http.ResponseWriter, *http.Request, *[]entities.Profile)
	PresentStoreMetadata(http.ResponseWriter, *http.Request, *entities.StoreMetadata)
	PresentListingMetadata(http.ResponseWriter, *http.Request, *entities.NonFungibleMetadata)
	PresentListings(http.ResponseWriter, *http.Request, *[]entities.Listing)
	PresentRefreshedProfile(w http.ResponseWriter, r *http.Request)
	PresentCurrency(w http.ResponseWriter, r *http.Request, currency *entities.Currency)
}

type presenter struct {
	logger   common.ILogger
	settings settings.ISettings
}

func (p *presenter) presentJSON(w http.ResponseWriter, r *http.Request, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	p.presentStatus(w, r, statusCode)
	json.NewEncoder(w).Encode(body)
}

func (p *presenter) presentText(w http.ResponseWriter, r *http.Request, statusCode int, text string) {
	w.Header().Set("Content-Type", "text/plain")
	p.presentStatus(w, r, statusCode)
	w.Write([]byte(text))
}

func (p *presenter) presentStatus(w http.ResponseWriter, r *http.Request, statusCode int) {
	p.logEvent(w, r, statusCode)
	w.WriteHeader(statusCode)
}

// log details of the request/response
func (p *presenter) logEvent(w http.ResponseWriter, r *http.Request, statusCode int) {
	ctx := r.Context()
	t1 := ctx.Value(common.ContextKeyRequestStartTime).(time.Time)
	p.logger.Info(ctx).Strs([]struct {
		Key   string
		Value string
	}{
		{Key: "method", Value: r.Method},
		{Key: "path", Value: r.URL.Path},
		{Key: "resptime", Value: time.Since(t1).String()},
		{Key: "statuscode", Value: fmt.Sprint(statusCode)},
		{Key: "remoteaddr", Value: r.RemoteAddr},
	}).Msg("http event")
}

func toChallengeJson(challenge *entities.Challenge) *challengeJson {
	return &challengeJson{
		Message: challenge.Message,
	}
}

func toProfileJson(profile *entities.Profile) *profileJson {
	return &profileJson{
		Address:      profile.Address,
		LastModified: profile.LastModified,
		ENSName:      profile.ENSName,
		StoreAssets: &storeAssetsJson{
			Premium:    profile.StoreAssets.Premium,
			BetaTester: profile.StoreAssets.BetaTester,
		},
		DisplayConfig:     toDisplayConfigJson(profile.DisplayConfig),
		NonFungibleTokens: toNonFungibleTokensJson(profile.NonFungibleTokens),
		FungibleTokens:    toFungibleTokensJson(profile.FungibleTokens),
		Statistics:        toStatisticsJson(profile.Statistics),
		Interactions:      toInteractionsJson(profile.Interactions),
		Currencies:        toCurrenciesJson(profile.Currencies),
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

func toCurrenciesJson(currencies *[]entities.Currency) *[]currencyJson {
	currenciesJson := []currencyJson{}

	for _, currency := range *currencies {
		currenciesJson = append(currenciesJson, *toCurrencyJson(&currency))
	}

	return &currenciesJson
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
	return &nonFungibleTokenJson{
		Contract: toContractJson(nft.Contract),
		TokenId:  nft.TokenId,
		Balance:  nft.Balance,
		Metadata: toNonFungibleMetadataJson(nft.Metadata),
	}
}

func toNonFungibleMetadataJson(metadata *entities.NonFungibleMetadata) *nonFungibleMetadataJson {
	if metadata == nil {
		return nil
	}

	return &nonFungibleMetadataJson{
		Name:        metadata.Name,
		Description: metadata.Description,
		Image:       metadata.Image,
		Attributes:  metadata.Attributes,
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

func toCurrencyJson(currency *entities.Currency) *currencyJson {
	return &currencyJson{
		Blockchain: currency.Blockchain,
		Balance:    currency.Balance,
	}
}

func toDisplayConfigJson(displayConfig *entities.DisplayConfig) *displayConfigJson {
	if displayConfig == nil {
		return nil
	}

	config := &displayConfigJson{
		Colors: &displayColorsJson{
			Primary:       displayConfig.Colors.Primary,
			Secondary:     displayConfig.Colors.Secondary,
			PrimaryText:   displayConfig.Colors.PrimaryText,
			SecondaryText: displayConfig.Colors.SecondaryText,
			Shadow:        displayConfig.Colors.Shadow,
			Accent:        displayConfig.Colors.Accent,
		},
		Info: &displayInfoJson{
			Title:         displayConfig.Info.Title,
			Description:   displayConfig.Info.Description,
			TwitterHandle: displayConfig.Info.TwitterHandle,
		},
		Picture: &displayPictureJson{},
		Achievements: &displayAchievementsJson{
			Text:  displayConfig.Achievements.Text,
			Items: &[]displayAchievementJson{},
		},
		Groups: &[]displayGroupJson{},
	}

	if displayConfig.Picture.Item != nil {
		config.Picture.Item = &displayItemJson{
			Id:    displayConfig.Picture.Item.Id,
			Index: displayConfig.Picture.Item.Index,
			Type:  displayConfig.Picture.Item.Type,
		}
	}

	for _, achievement := range *displayConfig.Achievements.Items {
		items := append(*config.Achievements.Items, displayAchievementJson{
			Id:    achievement.Id,
			Index: achievement.Index,
			Type:  achievement.Type,
		})
		config.Achievements.Items = &items
	}

	for _, group := range *displayConfig.Groups {
		groupJson := displayGroupJson{
			Id:    group.Id,
			Text:  group.Text,
			Items: &[]displayItemJson{},
		}

		for _, item := range *group.Items {
			items := append(*groupJson.Items, displayItemJson{
				Id:    item.Id,
				Index: item.Index,
				Type:  item.Type,
			})
			groupJson.Items = &items
		}

		groups := append(*config.Groups, groupJson)
		config.Groups = &groups
	}
	return config
}

func toStoreMetadataJson(metadata *entities.StoreMetadata) *storeMetadataJson {
	return &storeMetadataJson{
		Name:                 metadata.Name,
		Description:          metadata.Description,
		Image:                metadata.Image,
		ExternalLink:         metadata.ExternalLink,
		SellerFeeBasisPoints: metadata.SellerFeeBasisPoints,
		FeeRecipient:         metadata.FeeRecipient,
	}
}

func toListingsJson(listings *[]entities.Listing) *[]listingJson {
	listingJson := []listingJson{}

	for _, listing := range *listings {
		listingJson = append(listingJson, *toListingJson(&listing))
	}

	return &listingJson
}

func toListingJson(listing *entities.Listing) *listingJson {
	return &listingJson{
		Contract: toContractJson(listing.Contract),
		TokenId:  listing.TokenId,
		Info: &listingInfoJson{
			Purchasable:  listing.Info.Purchasable,
			Transferable: listing.Info.Transferable,
			Price:        listing.Info.Price,
			SupplyLimit:  listing.Info.SupplyLimit,
			BalanceLimit: listing.Info.BalanceLimit,
		},
		Metadata: toNonFungibleMetadataJson(listing.Metadata),
	}
}

type challengeJson struct {
	Message string `json:"message"`
}

type profileJson struct {
	Address           string                  `json:"address"`
	LastModified      *time.Time              `json:"last_modified,omitempty"`
	ENSName           string                  `json:"ens_name"`
	StoreAssets       *storeAssetsJson        `json:"store_assets"`
	DisplayConfig     *displayConfigJson      `json:"display_config,omitempty"`
	NonFungibleTokens *[]nonFungibleTokenJson `json:"non_fungible_tokens"`
	FungibleTokens    *[]fungibleTokenJson    `json:"fungible_tokens"`
	Statistics        *[]statisticJson        `json:"statistics"`
	Interactions      *[]interactionJson      `json:"interactions"`
	Currencies        *[]currencyJson         `json:"currencies"`
}

type contractJson struct {
	Blockchain common.Blockchain `json:"blockchain"`
	Address    string            `json:"address"`
	Interface  common.Interface  `json:"interface"`
}

type fungibleTokenJson struct {
	Contract *contractJson         `json:"contract"`
	Balance  *string               `json:"balance"`
	Metadata *fungibleMetadataJson `json:"metadata"`
}

type fungibleMetadataJson struct {
	Name     *string `json:"name"`
	Symbol   *string `json:"symbol"`
	Decimals *uint8  `json:"decimals"`
}

type nonFungibleTokenJson struct {
	Contract *contractJson            `json:"contract"`
	TokenId  string                   `json:"token_id"`
	Balance  *string                  `json:"balance"`
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

type currencyJson struct {
	common.Blockchain `json:"blockchain"`
	Balance           *string `json:"balance"`
}

type storeAssetsJson struct {
	Premium    bool `json:"premium"`
	BetaTester bool `json:"beta_tester"`
}

type displayConfigJson struct {
	Colors       *displayColorsJson       `json:"colors"`
	Info         *displayInfoJson         `json:"info"`
	Picture      *displayPictureJson      `json:"picture"`
	Achievements *displayAchievementsJson `json:"achievements"`
	Groups       *[]displayGroupJson      `json:"groups"`
}

type displayColorsJson struct {
	Primary       string `json:"primary"`
	Secondary     string `json:"secondary"`
	PrimaryText   string `json:"primary_text"`
	SecondaryText string `json:"secondary_text"`
	Shadow        string `json:"shadow"`
	Accent        string `json:"accent"`
}

type displayInfoJson struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	TwitterHandle string `json:"twitter_handle"`
}

type displayPictureJson struct {
	Item *displayItemJson `json:"item"` // Item can be nil
}

type displayAchievementsJson struct {
	Text  string                    `json:"text"`
	Items *[]displayAchievementJson `json:"items"`
}

type displayAchievementJson struct {
	Id    string                 `json:"id"`
	Index uint64                 `json:"index"`
	Type  common.AchievementType `json:"type"`
}

type displayGroupJson struct {
	Id    string             `json:"id"`
	Text  string             `json:"text"`
	Items *[]displayItemJson `json:"items"`
}

type displayItemJson struct {
	Id    string           `json:"id"`
	Index uint64           `json:"index"`
	Type  common.BadgeType `json:"type"`
}

type storeMetadataJson struct {
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Image                string `json:"image"`
	ExternalLink         string `json:"external_link"`
	SellerFeeBasisPoints uint   `json:"seller_fee_basis_points"`
	FeeRecipient         string `json:"fee_recipient"`
}

type listingJson struct {
	Contract *contractJson            `json:"contract"`
	TokenId  string                   `json:"token_id"`
	Info     *listingInfoJson         `json:"info"`
	Metadata *nonFungibleMetadataJson `json:"metadata"`
}

type listingInfoJson struct {
	Purchasable  bool   `json:"purchasable"`
	Transferable bool   `json:"transferable"`
	Price        string `json:"price"`
	BalanceLimit string `json:"balance_limit"`
	SupplyLimit  string `json:"supply_limit"`
}
