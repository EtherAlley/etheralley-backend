package http

import (
	"encoding/json"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

func (hc *HttpController) getProfileByAddressRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	address := ctx.Value(common.ContextKeyAddress).(string)

	profile, err := hc.getProfile(ctx, address)

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	hc.presenter.PresentProfile(ctx, w, profile)
}

func (hc *HttpController) saveProfileRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	address := ctx.Value(common.ContextKeyAddress).(string)

	profile := &entities.Profile{}
	err := json.NewDecoder(r.Body).Decode(profile)

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	err = hc.saveProfile(ctx, address, profile)

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	hc.presenter.PresentSavedProfile(ctx, w)
}
