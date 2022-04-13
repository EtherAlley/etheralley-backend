package http

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/entities"
)

func (p *httpPresenter) PresentCurrency(w http.ResponseWriter, r *http.Request, currency *entities.Currency) {
	json := toCurrencyJson(currency)

	p.presentJSON(w, r, http.StatusOK, json)
}
