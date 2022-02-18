package http

import (
	"encoding/json"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/usecases"
)

func (hc *HttpController) getProfileByAddressRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	address := ctx.Value(common.ContextKeyAddress).(string)

	profile, err := hc.getProfile(ctx, &usecases.GetProfileInput{
		Address: address,
	})

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	hc.presenter.PresentProfile(ctx, w, profile)
}

func (hc *HttpController) saveProfileRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profile := &usecases.ProfileInput{}
	err := json.NewDecoder(r.Body).Decode(profile)

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	err = hc.saveProfile(ctx, &usecases.SaveProfileInput{
		Profile: profile,
	})

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	hc.presenter.PresentSavedProfile(ctx, w)
}
