package http

import (
	"context"
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

func (p *httpPresenter) PresentFungibleToken(ctx context.Context, w http.ResponseWriter, token *entities.FungibleToken) {
	json := toFungibleJson(token)

	render(w, http.StatusOK, json)
}

func (p *httpPresenter) PresentNonFungibleToken(ctx context.Context, w http.ResponseWriter, nft *entities.NonFungibleToken) {
	json := toNonFungibleJson(nft)

	render(w, http.StatusOK, json)
}

func (p *httpPresenter) PresentStatistic(ctx context.Context, w http.ResponseWriter, stat *entities.Statistic) {

	json := toStatisticJson(stat)

	if json != nil {
		render(w, http.StatusOK, json)
		return
	}
}
