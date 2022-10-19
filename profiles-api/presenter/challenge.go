package presenter

import (
	"net/http"

	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

func (p *presenter) PresentChallenge(w http.ResponseWriter, r *http.Request, challenge *entities.Challenge) {
	json := toChallengeJson(challenge)
	p.presentJSON(w, r, http.StatusOK, json)
}
