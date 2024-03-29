package controller

import (
	"net/http"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/usecases"
)

func (hc *controller) getChallengeRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	address := ctx.Value(common.ContextKeyAddress).(string)

	challenge, err := hc.getChallenge.Do(ctx, &usecases.GetChallengeInput{
		Address: address,
	})

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentChallenge(w, r, challenge)
}
