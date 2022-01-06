package http

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) registerNFTRoutes(r chi.Router) {
	r.Get("/", hc.getNFT)
}

func (hc *HttpController) getNFT(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	address := r.Context().Value(contextKeyAddress).(string)
	location := &entities.NFTLocation{
		ContractAddress: query.Get("contract_address"),
		SchemaName:      query.Get("schema_name"),
		TokenId:         query.Get("token_id"),
		Blockchain:      query.Get("blockchain"),
	}

	nft, err := hc.getNFTUseCase(r.Context(), address, location)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, err)
		return
	}

	Render(w, http.StatusOK, nft)
}
