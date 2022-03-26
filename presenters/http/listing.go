package http

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

func (p *httpPresenter) PresentListingMetadata(w http.ResponseWriter, r *http.Request, metadata *entities.NonFungibleMetadata) {
	json := toNonFungibleMetadataJson(metadata)

	p.presentJSON(w, r, http.StatusOK, json)
}

func (p *httpPresenter) PresentListings(w http.ResponseWriter, r *http.Request, listings *[]entities.Listing) {
	json := toListingsJson(listings)

	p.presentJSON(w, r, http.StatusOK, json)
}
