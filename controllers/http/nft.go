package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) registerNFTRoutes(r chi.Router) {
	r.Get("/", hc.getNFT)
}

func (hc *HttpController) getNFT(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	address := query.Get("address")
	blockchain := query.Get("blockchain")
	contractAddress := query.Get("contract_address")
	schemaName := query.Get("schema_name")
	tokenId := query.Get("token_id")

	nft, err := hc.getNFTUseCase(r.Context(), address, blockchain, contractAddress, schemaName, tokenId)

	if err != nil {
		RenderErr(w, http.StatusBadRequest, err)
		return
	}

	Render(w, http.StatusOK, nft)
}
