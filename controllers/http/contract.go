package http

import (
	"context"
	"net/http"

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

		ctx = context.WithValue(ctx, contextKeyContract, contract)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (hc *HttpController) getTokenRoute(w http.ResponseWriter, r *http.Request) {
	contract := r.Context().Value(contextKeyContract).(*entities.Contract)

	query := r.URL.Query()
	address := query.Get("user_address")

	token, err := hc.getFungibleToken(r.Context(), address, contract)

	if err != nil {
		RenderError(w, http.StatusBadRequest, "bad request")
		return
	}

	Render(w, http.StatusOK, token)
}

func (hc *HttpController) getNFTRoute(w http.ResponseWriter, r *http.Request) {
	contract := r.Context().Value(contextKeyContract).(*entities.Contract)

	query := r.URL.Query()
	address := query.Get("user_address")
	tokenId := query.Get("token_id")

	nft, err := hc.getNonFungibleToken(r.Context(), address, contract, tokenId)

	if err != nil {
		RenderError(w, http.StatusBadRequest, "bad request")
		return
	}

	Render(w, http.StatusOK, nft)
}

func (hc *HttpController) getStatisticRoute(w http.ResponseWriter, r *http.Request) {
	contract := r.Context().Value(contextKeyContract).(*entities.Contract)

	query := r.URL.Query()
	address := query.Get("user_address")
	statType := query.Get("type")

	statistic, err := hc.getStatistic(r.Context(), address, contract, statType)

	if err != nil {
		RenderError(w, http.StatusBadRequest, "bad request")
		return
	}

	Render(w, http.StatusOK, statistic)
}
