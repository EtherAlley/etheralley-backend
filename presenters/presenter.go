package presenters

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

type IPresenter interface {
	PresentBadRequest(http.ResponseWriter, *http.Request, error)
	PresentUnathorized(http.ResponseWriter, *http.Request)
	PresentHealth(http.ResponseWriter, *http.Request)
	PresentChallenge(http.ResponseWriter, *http.Request, *entities.Challenge)
	PresentFungibleToken(http.ResponseWriter, *http.Request, *entities.FungibleToken)
	PresentNonFungibleToken(http.ResponseWriter, *http.Request, *entities.NonFungibleToken)
	PresentStatistic(http.ResponseWriter, *http.Request, *entities.Statistic)
	PresentInteraction(http.ResponseWriter, *http.Request, *entities.Interaction)
	PresentProfile(http.ResponseWriter, *http.Request, *entities.Profile)
	PresentSavedProfile(http.ResponseWriter, *http.Request)
	PresentTopProfiles(http.ResponseWriter, *http.Request, *[]entities.Profile)
}
