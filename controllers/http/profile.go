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
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentProfile(w, r, profile)
}

func (hc *HttpController) saveProfileRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profile := &usecases.ProfileInput{}
	err := json.NewDecoder(r.Body).Decode(profile)

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	profile.Address = ctx.Value(common.ContextKeyAddress).(string)

	err = hc.saveProfile(ctx, &usecases.SaveProfileInput{
		Profile: profile,
	})

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentSavedProfile(w, r)
}

func (hc *HttpController) recordProfileViewMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// we don't need to block for this.
		done := make(chan bool)
		go func() {
			// we also don't care about the results, they will not affect the results of the request
			hc.recordProfileView(ctx, &usecases.RecordProfileViewInput{
				Address:   ctx.Value(common.ContextKeyAddress).(string),
				IpAddress: r.RemoteAddr,
			})
			done <- true
		}()

		next.ServeHTTP(w, r)

		<-done
	})
}

func (hc *HttpController) getTopProfilesRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profiles := hc.getTopProfiles(ctx, &usecases.GetTopProfilesInput{})

	hc.presenter.PresentTopProfiles(w, r, profiles)
}
