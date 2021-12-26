package http

import (
	"encoding/json"
	"net/http"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/entities"
	"github.com/eflem00/go-example-app/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func (handler *ProfileHandler) GetProfileByAddress(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	profile, err := handler.profileUsecase.GetProfileByAddress(r.Context(), address)

	if err == mongo.ErrNoDocuments {
		render.Render(w, r, &ErrResponse{
			Err:            err,
			HTTPStatusCode: 404,
			StatusText:     "Not found.",
			ErrorText:      err.Error(),
		})
		return
	}

	if err != nil {
		render.Render(w, r, &ErrResponse{
			Err:            err,
			HTTPStatusCode: 400,
			StatusText:     "Invalid request.",
			ErrorText:      err.Error(),
		})
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
		render.Render(w, r, &ErrResponse{
			Err:            err,
			HTTPStatusCode: 400,
			StatusText:     "Invalid request.",
			ErrorText:      err.Error(),
		})
		return
	}

	err = handler.profileUsecase.SaveProfile(r.Context(), address, profile)

	if err != nil {
		render.Render(w, r, &ErrResponse{
			Err:            err,
			HTTPStatusCode: 400,
			StatusText:     "Invalid request.",
			ErrorText:      err.Error(),
		})
		return
	}

	render.Status(r, http.StatusCreated)
}
