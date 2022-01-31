package http

import (
	"encoding/json"
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

func (hc *HttpController) getProfileByAddressRoute(w http.ResponseWriter, r *http.Request) {
	address := r.Context().Value(contextKeyAddress).(string)

	profile, err := hc.getProfile(r.Context(), address)

	if err != nil {
		RenderError(w, http.StatusBadRequest, "bad request")
		return
	}

	Render(w, http.StatusOK, profile)
}

func (hc *HttpController) saveProfileRoute(w http.ResponseWriter, r *http.Request) {
	address := r.Context().Value(contextKeyAddress).(string)

	profile := &entities.Profile{}
	err := json.NewDecoder(r.Body).Decode(profile)

	if err != nil {
		RenderError(w, http.StatusBadRequest, "bad request")
		return
	}

	profile.Address = address

	err = hc.saveProfile(r.Context(), profile)

	if err != nil {
		RenderError(w, http.StatusBadRequest, "bad request")
		return
	}

	RenderNoBody(w, http.StatusCreated)
}
