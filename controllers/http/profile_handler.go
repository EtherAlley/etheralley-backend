package http

import (
	"encoding/json"
	"net/http"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/entities"
	"github.com/eflem00/go-example-app/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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

type ResponseError struct {
	Message string `json:"message"`
}

func (handler *ProfileHandler) GetProfileByAddress(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	profile, err := handler.profileUsecase.GetProfileByAddress(r.Context(), address)

	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ResponseError{"Error fetching profile"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, profile)
}

func (handler *ProfileHandler) SaveProfile(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	var profile entities.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)

	if err != nil {
		handler.logger.Err(err, "error decoding json")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ResponseError{"Error saving profile"})
		return
	}

	err = handler.profileUsecase.SaveProfile(r.Context(), address, profile)

	if err != nil {
		handler.logger.Err(err, "error fetching profile")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ResponseError{"Error saving profile"})
		return
	}

	render.Status(r, http.StatusCreated)
}
