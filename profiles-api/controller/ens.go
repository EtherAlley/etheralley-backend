package controller

import (
	"context"
	"net/http"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/usecases"
	"github.com/go-chi/chi/v5"
)

func (hc *controller) resolveAddressRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// the address param of the route could be either an ens name or an address
		address, err := hc.resolveAddress.Do(ctx, &usecases.ResolveAddressInput{
			Value: chi.URLParam(r, "address"),
		})

		if err != nil {
			hc.presenter.PresentNotFound(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, common.ContextKeyAddress, address)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
