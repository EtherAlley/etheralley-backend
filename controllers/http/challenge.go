package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) registerChallengeRoutes(r chi.Router) {
	r.Get("/", hc.getChallenge)
}

func (hc *HttpController) getChallenge(w http.ResponseWriter, r *http.Request) {
	address := r.Context().Value(contextKeyAddress).(string)

	challenge, err := hc.getChallengeUseCase(r.Context(), address)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, err)
		return
	}

	Render(w, http.StatusOK, challenge)
}
