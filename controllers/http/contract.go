package http

import (
	"context"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/usecases"
)

func (hc *HttpController) parseContract(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		query := r.URL.Query()

		contract := &usecases.ContractInput{
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
	query := r.URL.Query()

	token, err := hc.getFungibleToken(ctx, &usecases.GetFungibleTokenInput{
		Address: query.Get("user_address"),
		Token: &usecases.FungibleTokenInput{
			Contract: ctx.Value(common.ContextKeyContract).(*usecases.ContractInput),
		},
	})

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	hc.presenter.PresentFungibleToken(ctx, w, token)
}

func (hc *HttpController) getNFTRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()

	nft, err := hc.getNonFungibleToken(ctx, &usecases.GetNonFungibleTokenInput{
		Address: query.Get("user_address"),
		NonFungibleToken: &usecases.NonFungibleTokenInput{
			TokenId:  query.Get("token_id"),
			Contract: ctx.Value(common.ContextKeyContract).(*usecases.ContractInput),
		},
	})

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	hc.presenter.PresentNonFungibleToken(ctx, w, nft)
}

func (hc *HttpController) getStatisticRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()

	statistic, err := hc.getStatistic(ctx, &usecases.GetStatisticsInput{
		Address: query.Get("user_address"),
		Statistic: &usecases.StatisticInput{
			Contract: ctx.Value(common.ContextKeyContract).(*usecases.ContractInput),
			Type:     query.Get("type"),
		},
	})

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	hc.presenter.PresentStatistic(ctx, w, statistic)
}
