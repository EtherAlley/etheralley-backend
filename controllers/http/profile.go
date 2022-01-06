package http

import (
	"encoding/json"
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) registerProfileRoutes(r chi.Router) {
	r.Get("/", hc.getProfileByAddress)
	r.With(hc.authenticate).Put("/", hc.saveProfile)
}

func (hc *HttpController) getProfileByAddress(w http.ResponseWriter, r *http.Request) {
	address := r.Context().Value(contextKeyAddress).(string)

	profile, err := hc.getProfileUseCase(r.Context(), address)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, err)
		return
	}

	Render(w, http.StatusOK, profile)
}

func (hc *HttpController) saveProfile(w http.ResponseWriter, r *http.Request) {
	address := r.Context().Value(contextKeyAddress).(string)

	profile := &entities.Profile{}
	err := json.NewDecoder(r.Body).Decode(profile)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, err)
		return
	}

	profile.Address = address

	err = hc.saveProfileUseCase(r.Context(), profile)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, err)
		return
	}

	RenderNoBody(w, http.StatusCreated)
}
