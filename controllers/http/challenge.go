package http

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
)

func (hc *HttpController) getChallengeRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	address := ctx.Value(common.ContextKeyAddress).(string)

	challenge, err := hc.getChallenge(ctx, address)

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	hc.presenter.PresentChallenge(ctx, w, challenge)
}
