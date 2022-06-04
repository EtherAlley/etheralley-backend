package controller

import (
	"net/http"

	"github.com/etheralley/etheralley-apis/core/usecases"
)

func (hc *controller) getCurrencyRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()

	currency, err := hc.getCurrency.Do(ctx, &usecases.GetCurrencyInput{
		Address:    query.Get("address"),
		Blockchain: query.Get("blockchain"),
	})

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentCurrency(w, r, currency)
}
