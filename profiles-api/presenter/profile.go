package presenter

import (
	"net/http"

	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

func (p *presenter) PresentProfile(w http.ResponseWriter, r *http.Request, profile *entities.Profile) {
	json := toProfileJson(profile)

	p.presentJSON(w, r, http.StatusOK, json)
}

func (p *presenter) PresentSavedProfile(w http.ResponseWriter, r *http.Request) {
	p.presentStatus(w, r, http.StatusCreated)
}

func (p *presenter) PresentProfiles(w http.ResponseWriter, r *http.Request, profiles *[]entities.Profile) {
	json := []profileJson{}

	for _, profile := range *profiles {
		json = append(json, *toProfileJson(&profile))
	}

	p.presentJSON(w, r, http.StatusOK, json)
}

func (p *presenter) PresentRefreshedProfile(w http.ResponseWriter, r *http.Request) {
	p.presentStatus(w, r, http.StatusOK)
}
