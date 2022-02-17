package http

import (
	"context"
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

func (p *httpPresenter) PresentInteraction(ctx context.Context, w http.ResponseWriter, interaction *entities.Interaction) {
	json := toInteractionJson(interaction)

	render(w, http.StatusOK, json)
}
