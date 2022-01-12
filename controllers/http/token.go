package http

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) registerTokenRoutes(r chi.Router) {
	r.Get("/", hc.getTokenRoute)
}

func (hc *HttpController) getTokenRoute(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	address := r.Context().Value(contextKeyAddress).(string)
	contract := &entities.Contract{
		Address:    query.Get("address"),
		Interface:  query.Get("interface"),
		Blockchain: query.Get("blockchain"),
	}

	if contract.Interface == common.ERC20 {
		token, err := hc.getFungibleToken(r.Context(), address, contract)

		if err != nil {
			RenderErr(w, http.StatusBadRequest, err)
			return
		}

		Render(w, http.StatusOK, token)
		return
	}

	tokenId := query.Get("token_id")

	nft, err := hc.getNonFungibleToken(r.Context(), address, contract, tokenId)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, err)
		return
	}

	Render(w, http.StatusOK, nft)
}
