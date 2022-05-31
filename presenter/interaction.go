package presenter

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

func (p *presenter) PresentInteraction(w http.ResponseWriter, r *http.Request, interaction *entities.Interaction) {
	json := toInteractionJson(interaction)

	p.presentJSON(w, r, http.StatusOK, json)
}
