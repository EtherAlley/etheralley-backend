package http

import (
	"context"
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

func (p *httpPresenter) PresentChallenge(ctx context.Context, w http.ResponseWriter, challenge *entities.Challenge) {
	json := toChallengeJson(challenge)
	render(w, http.StatusOK, json)
}

func toChallengeJson(challenge *entities.Challenge) *challengeJson {
	return &challengeJson{
		Message: challenge.Message,
	}
}

type challengeJson struct {
	Message string `json:"message"`
}
