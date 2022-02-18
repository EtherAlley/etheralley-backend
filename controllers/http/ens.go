package http

import (
	"context"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/usecases"
	"github.com/go-chi/chi/v5"
)

// the address param of the route could be either an ens name or an address
func (hc *HttpController) resolveAddr(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		address, err := hc.resolveAddress(ctx, &usecases.ResolveAddressInput{
			Value: chi.URLParam(r, "address"),
		})

		if err != nil {
			hc.presenter.PresentBadRequest(ctx, w, err)
			return
		}

		ctx = context.WithValue(ctx, common.ContextKeyAddress, address)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
