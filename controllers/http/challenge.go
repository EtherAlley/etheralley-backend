package http

import (
	"net/http"
)

func (hc *HttpController) getChallengeRoute(w http.ResponseWriter, r *http.Request) {
	address := r.Context().Value(contextKeyAddress).(string)

	challenge, err := hc.getChallenge(r.Context(), address)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, err)
		return
	}

	Render(w, http.StatusOK, challenge)
}
