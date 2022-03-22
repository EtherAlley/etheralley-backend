package http

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

func (p *httpPresenter) PresentFungibleToken(w http.ResponseWriter, r *http.Request, token *entities.FungibleToken) {
	json := toFungibleJson(token)

	p.presentJSON(w, r, http.StatusOK, json)
}

func (p *httpPresenter) PresentNonFungibleToken(w http.ResponseWriter, r *http.Request, nft *entities.NonFungibleToken) {
	json := toNonFungibleJson(nft)

	p.presentJSON(w, r, http.StatusOK, json)
}

func (p *httpPresenter) PresentStatistic(w http.ResponseWriter, r *http.Request, stat *entities.Statistic) {

	json := toStatisticJson(stat)

	if json != nil {
		p.presentJSON(w, r, http.StatusOK, json)
		return
	}
}
