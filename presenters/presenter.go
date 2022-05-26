package presenters

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

type IPresenter interface {
	PresentBadRequest(http.ResponseWriter, *http.Request, error)
	PresentUnathorized(http.ResponseWriter, *http.Request, error)
	PresentNotFound(http.ResponseWriter, *http.Request, error)
	PresentTooManyRequests(http.ResponseWriter, *http.Request, error)
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
