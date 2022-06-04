package presenter

import (
	"net/http"

	"github.com/etheralley/etheralley-backend/core/entities"
)

func (p *presenter) PresentStoreMetadata(w http.ResponseWriter, r *http.Request, metadata *entities.StoreMetadata) {
	json := toStoreMetadataJson(metadata)

	p.presentJSON(w, r, http.StatusOK, json)
}

func (p *presenter) PresentListingMetadata(w http.ResponseWriter, r *http.Request, metadata *entities.NonFungibleMetadata) {
	json := toNonFungibleMetadataJson(metadata)

	p.presentJSON(w, r, http.StatusOK, json)
}

func (p *presenter) PresentListings(w http.ResponseWriter, r *http.Request, listings *[]entities.Listing) {
	json := toListingsJson(listings)

	p.presentJSON(w, r, http.StatusOK, json)
}
