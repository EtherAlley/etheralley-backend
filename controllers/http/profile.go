package http

import (
	"encoding/json"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) registerProfileRoutes(r chi.Router) {
	r.Get("/{address}", hc.getProfileByAddress)
	r.With(hc.authenticate).Put("/{address}", hc.saveProfile)
}

func (hc *HttpController) getProfileByAddress(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	profile, err := hc.getProfileUsecase.Go(r.Context(), address)

	if err == common.ErrNil {
		RenderNoBody(w, http.StatusNotFound)
		return
	}

	if err != nil {
		RenderNoBody(w, http.StatusBadRequest)
		return
	}

	Render(w, http.StatusOK, profile)
}

func (hc *HttpController) saveProfile(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	profile := &entities.Profile{}
	err := json.NewDecoder(r.Body).Decode(profile)

	if err != nil {
		RenderNoBody(w, http.StatusBadRequest)
		return
	}

	profile.Address = address

	err = hc.saveProfileUsecase.Go(r.Context(), profile)

	if err != nil {
		RenderNoBody(w, http.StatusBadRequest)
		return
	}

	RenderNoBody(w, http.StatusCreated)
}
