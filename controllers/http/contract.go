package http

import (
	"context"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

func (hc *HttpController) parseContract(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		ctx := r.Context()

		contract := &entities.Contract{
			Address:    query.Get("contract"),
			Interface:  query.Get("interface"),
			Blockchain: query.Get("blockchain"),
		}

		ctx = context.WithValue(ctx, common.ContextKeyContract, contract)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (hc *HttpController) getTokenRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contract := ctx.Value(common.ContextKeyContract).(*entities.Contract)

	query := r.URL.Query()
	address := query.Get("user_address")

	token, err := hc.getFungibleToken(ctx, address, contract)

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	hc.presenter.PresentFungibleToken(ctx, w, token)
}

func (hc *HttpController) getNFTRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contract := ctx.Value(common.ContextKeyContract).(*entities.Contract)

	query := r.URL.Query()
	address := query.Get("user_address")
	tokenId := query.Get("token_id")

	nft, err := hc.getNonFungibleToken(ctx, address, contract, tokenId)

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	hc.presenter.PresentNonFungibleToken(ctx, w, nft)
}

func (hc *HttpController) getStatisticRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contract := ctx.Value(common.ContextKeyContract).(*entities.Contract)

	query := r.URL.Query()
	address := query.Get("user_address")
	statType := query.Get("type")

	statistic, err := hc.getStatistic(ctx, address, contract, statType)

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	hc.presenter.PresentStatistic(ctx, w, statistic)
}
