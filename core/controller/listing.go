package controller

import (
	"encoding/json"
	"net/http"

	"github.com/etheralley/etheralley-apis/core/usecases"
	"github.com/go-chi/chi/v5"
)

func (hc *controller) getStoreMetadataRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	metadata := hc.getStoreMetadata.Do(ctx)

	hc.presenter.PresentStoreMetadata(w, r, metadata)
}

func (hc *controller) getMetadataByIdRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	metadata, err := hc.getListingMetadata.Do(ctx, &usecases.GetListingMetadataInput{
		TokenId: chi.URLParam(r, "tokenid"),
	})

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentListingMetadata(w, r, metadata)
}

func (hc *controller) getListingsRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := &usecases.GetListingsInput{}
	err := json.NewDecoder(r.Body).Decode(input)

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	listings, err := hc.getListings.Do(ctx, input)

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentListings(w, r, listings)
}
