package http

import (
	"net/http"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/usecases"
	"github.com/go-chi/chi/v5"
)

type ChallengeHandler struct {
	logger      *common.Logger
	authUseCase *usecases.AuthenticationUseCase
}

func NewChallengeHandler(authUseCase *usecases.AuthenticationUseCase, logger *common.Logger) *ChallengeHandler {
	return &ChallengeHandler{
		logger,
		authUseCase,
	}
}

func (handler *ChallengeHandler) GetChallenge(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	challenge, err := handler.authUseCase.GetChallenge(r.Context(), address)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, "Invalid request.")
		return
	}

	Render(w, http.StatusOK, challenge)
}
