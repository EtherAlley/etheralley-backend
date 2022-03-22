package http

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

func (p *httpPresenter) PresentChallenge(w http.ResponseWriter, r *http.Request, challenge *entities.Challenge) {
	json := toChallengeJson(challenge)
	p.presentJSON(w, r, http.StatusOK, json)
}
