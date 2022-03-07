package presenters

import (
	"context"
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

type IPresenter interface {
	PresentBadRequest(context.Context, http.ResponseWriter, error)
	PresentUnathorized(context.Context, http.ResponseWriter)
	PresentHealth(context.Context, http.ResponseWriter)
	PresentChallenge(context.Context, http.ResponseWriter, *entities.Challenge)
	PresentFungibleToken(context.Context, http.ResponseWriter, *entities.FungibleToken)
	PresentNonFungibleToken(context.Context, http.ResponseWriter, *entities.NonFungibleToken)
	PresentStatistic(context.Context, http.ResponseWriter, *entities.Statistic)
	PresentInteraction(context.Context, http.ResponseWriter, *entities.Interaction)
	PresentProfile(context.Context, http.ResponseWriter, *entities.Profile)
	PresentSavedProfile(context.Context, http.ResponseWriter)
	PresentTopProfiles(ctx context.Context, w http.ResponseWriter, profiles *[]entities.Profile)
}
