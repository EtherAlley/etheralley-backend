package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) registerChallengeRoutes(r chi.Router) {
	r.Get("/{address}", hc.getChallenge)
}

func (hc *HttpController) getChallenge(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	challenge, err := hc.getChallengeUsecase(r.Context(), address)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, err)
		return
	}

	Render(w, http.StatusOK, challenge)
}
