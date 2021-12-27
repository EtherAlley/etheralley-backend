package http

import (
	"net/http"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/usecases"
	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	logger      *common.Logger
	authUseCase *usecases.AuthenticationUseCase
}

func NewAuthHandler(authUseCase *usecases.AuthenticationUseCase, logger *common.Logger) *AuthHandler {
	return &AuthHandler{
		logger,
		authUseCase,
	}
}

type ChallengeMessageBody struct {
	Message string `json:"message"`
}

func (handler *AuthHandler) GetChallengeMessage(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	msg, err := handler.authUseCase.GetChallengeMessage(r.Context(), address)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, "Invalid request.")
		return
	}

	Render(w, http.StatusOK, ChallengeMessageBody{Message: msg})
}
