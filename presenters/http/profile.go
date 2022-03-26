package http

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

func (p *httpPresenter) PresentProfile(w http.ResponseWriter, r *http.Request, profile *entities.Profile) {
	json := toProfileJson(profile)

	p.presentJSON(w, r, http.StatusOK, json)
}

func (p *httpPresenter) PresentSavedProfile(w http.ResponseWriter, r *http.Request) {
	p.presentStatus(w, r, http.StatusCreated)
}

func (p *httpPresenter) PresentTopProfiles(w http.ResponseWriter, r *http.Request, profiles *[]entities.Profile) {
	json := []profileJson{}

	for _, profile := range *profiles {
		json = append(json, *toProfileJson(&profile))
	}

	p.presentJSON(w, r, http.StatusOK, json)
}

func (p *httpPresenter) PresentRefreshedProfile(w http.ResponseWriter, r *http.Request) {
	p.presentStatus(w, r, http.StatusOK)
}
