package http

import (
	"encoding/json"
	"net/http"

	"github.com/etheralley/etheralley-core-api/usecases"
	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) getMetadataByIdRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	metadata, err := hc.getListingMetadata(ctx, &usecases.GetListingMetadataInput{
		TokenId: chi.URLParam(r, "tokenid"),
	})

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentListingMetadata(w, r, metadata)
}

func (hc *HttpController) getListingsRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := &usecases.GetListingsInput{}
	err := json.NewDecoder(r.Body).Decode(input)

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	listings, err := hc.getListings(ctx, input)

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentListings(w, r, listings)
}
