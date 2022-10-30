package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
	"github.com/etheralley/etheralley-backend/profiles-api/usecases"
)

func (hc *controller) getProfileByAddressRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	address := ctx.Value(common.ContextKeyAddress).(string)
	hydration := r.URL.Query().Get("hydration")

	var profile *entities.Profile
	var err error
	if hydration == common.LIGHT {
		profile, err = hc.getLightProfile.Do(ctx, &usecases.GetLightProfileInput{Address: address})
	} else {
		profile, err = hc.getProfile.Do(ctx, &usecases.GetProfileInput{
			Address: address,
		})
	}

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	if profile.Banned {
		hc.presenter.PresentForbiddenRequest(w, r, fmt.Errorf("banned profile"))
		return
	}

	hc.presenter.PresentProfile(w, r, profile)
}

func (hc *controller) saveProfileRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profile := &usecases.ProfileInput{}
	err := json.NewDecoder(r.Body).Decode(profile)

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	profile.Address = ctx.Value(common.ContextKeyAddress).(string)

	err = hc.saveProfile.Do(ctx, &usecases.SaveProfileInput{
		Profile: profile,
	})

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentSavedProfile(w, r)
}

func (hc *controller) recordProfileViewMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// we don't need to block for this.
		done := make(chan bool)
		go func() {
			// we also don't care about the results, they will not affect the results of the request
			hc.recordProfileView.Do(ctx, &usecases.RecordProfileViewInput{
				Address:   ctx.Value(common.ContextKeyAddress).(string),
				IpAddress: r.RemoteAddr,
			})
			done <- true
		}()

		next.ServeHTTP(w, r)

		<-done
	})
}

func (hc *controller) getTopProfilesRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profiles := hc.getTopProfiles.Do(ctx, &usecases.GetTopProfilesInput{})

	hc.presenter.PresentProfiles(w, r, profiles)
}

func (hc *controller) getSpotlightProfileRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profiles := hc.getSpotlightProfiles.Do(ctx, &usecases.GetSpotlightProfilesInput{})

	hc.presenter.PresentProfiles(w, r, profiles)
}

func (hc *controller) refreshProfileRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	address := ctx.Value(common.ContextKeyAddress).(string)

	err := hc.refreshProfile.Do(ctx, &usecases.RefreshProfileInput{
		Address: address,
	})

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentRefreshedProfile(w, r)
}
