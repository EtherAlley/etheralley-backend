package controller

import (
	"context"
	"net/http"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/usecases"
)

func (hc *controller) parseContract(next http.Handler) http.Handler {
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

func (hc *controller) getTokenRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()

	token, err := hc.getFungibleToken.Do(ctx, &usecases.GetFungibleTokenInput{
		Address: query.Get("user_address"),
		Token: &usecases.FungibleTokenInput{
			Contract: ctx.Value(common.ContextKeyContract).(*usecases.ContractInput),
		},
	})

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentFungibleToken(w, r, token)
}

func (hc *controller) getNFTRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()

	nft, err := hc.getNonFungibleToken.Do(ctx, &usecases.GetNonFungibleTokenInput{
		Address: query.Get("user_address"),
		NonFungibleToken: &usecases.NonFungibleTokenInput{
			TokenId:  query.Get("token_id"),
			Contract: ctx.Value(common.ContextKeyContract).(*usecases.ContractInput),
		},
	})

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentNonFungibleToken(w, r, nft)
}

func (hc *controller) getStatisticRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()

	statistic, err := hc.getStatistic.Do(ctx, &usecases.GetStatisticsInput{
		Address: query.Get("user_address"),
		Statistic: &usecases.StatisticInput{
			Contract: ctx.Value(common.ContextKeyContract).(*usecases.ContractInput),
			Type:     query.Get("type"),
		},
	})

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentStatistic(w, r, statistic)
}
