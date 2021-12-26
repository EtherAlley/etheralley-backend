package http

import (
	"encoding/json"
	"net/http"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/entities"
	"github.com/eflem00/go-example-app/usecases"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileHandler struct {
	profileUsecase *usecases.ProfileUsecase
	logger         *common.Logger
}

func NewProfileHandler(profileUsecase *usecases.ProfileUsecase, logger *common.Logger) *ProfileHandler {
	return &ProfileHandler{
		profileUsecase,
		logger,
	}
}

func (handler *ProfileHandler) GetProfileByAddress(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	profile, err := handler.profileUsecase.GetProfileByAddress(r.Context(), address)

	if err == mongo.ErrNoDocuments {
		RenderErr(w, http.StatusNotFound, "Not found.")
		return
	}

	if err != nil {
		RenderErr(w, http.StatusBadRequest, "Invalid request.")
		return
	}

	Render(w, http.StatusOK, profile)
}

func (handler *ProfileHandler) SaveProfile(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	var profile entities.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, "Invalid request.")
		return
	}

	err = handler.profileUsecase.SaveProfile(r.Context(), address, profile)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, "Invalid request.")
		return
	}

	RenderNoBody(w, http.StatusCreated)
}
