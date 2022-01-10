package http

import (
	"encoding/json"
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) registerProfileRoutes(r chi.Router) {
	r.Get("/", hc.getProfileByAddressRoute)
	r.With(hc.authenticate).Put("/", hc.saveProfileRoute)
}

func (hc *HttpController) getProfileByAddressRoute(w http.ResponseWriter, r *http.Request) {
	address := r.Context().Value(contextKeyAddress).(string)

	profile, err := hc.getProfile(r.Context(), address)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, err)
		return
	}

	Render(w, http.StatusOK, profile)
}

func (hc *HttpController) saveProfileRoute(w http.ResponseWriter, r *http.Request) {
	address := r.Context().Value(contextKeyAddress).(string)

	profile := &entities.Profile{}
	err := json.NewDecoder(r.Body).Decode(profile)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, err)
		return
	}

	profile.Address = address

	err = hc.saveProfile(r.Context(), profile)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, err)
		return
	}

	RenderNoBody(w, http.StatusCreated)
}
