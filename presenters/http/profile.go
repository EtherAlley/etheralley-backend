package http

import (
	"context"
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

func (p *httpPresenter) PresentProfile(ctx context.Context, w http.ResponseWriter, profile *entities.Profile) {
	json := toProfileJson(profile)

	render(w, http.StatusOK, json)
}

func (p *httpPresenter) PresentSavedProfile(ctx context.Context, w http.ResponseWriter) {
	renderNoBody(w, http.StatusCreated)
}
