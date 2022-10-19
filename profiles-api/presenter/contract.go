package presenter

import (
	"net/http"

	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

func (p *presenter) PresentFungibleToken(w http.ResponseWriter, r *http.Request, token *entities.FungibleToken) {
	json := toFungibleJson(token)

	p.presentJSON(w, r, http.StatusOK, json)
}

func (p *presenter) PresentNonFungibleToken(w http.ResponseWriter, r *http.Request, nft *entities.NonFungibleToken) {
	json := toNonFungibleJson(nft)

	p.presentJSON(w, r, http.StatusOK, json)
}

func (p *presenter) PresentStatistic(w http.ResponseWriter, r *http.Request, stat *entities.Statistic) {

	json := toStatisticJson(stat)

	if json != nil {
		p.presentJSON(w, r, http.StatusOK, json)
		return
	}
}
