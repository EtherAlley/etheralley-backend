package http

import (
	"net/http"
)

func (hc *HttpController) getChallengeRoute(w http.ResponseWriter, r *http.Request) {
	address := r.Context().Value(contextKeyAddress).(string)

	challenge, err := hc.getChallenge(r.Context(), address)

	if err != nil {
		RenderError(w, http.StatusBadRequest, "bad request")
		return
	}

	Render(w, http.StatusOK, challenge)
}
