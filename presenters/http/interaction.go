package http

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

func (p *httpPresenter) PresentInteraction(w http.ResponseWriter, r *http.Request, interaction *entities.Interaction) {
	json := toInteractionJson(interaction)

	p.presentJSON(w, r, http.StatusOK, json)
}
