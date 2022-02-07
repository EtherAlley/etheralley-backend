package http

import (
	"context"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/go-chi/chi/v5"
)

// the address param of the route could be either an ens name or an address
func (hc *HttpController) resolveAddr(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input := chi.URLParam(r, "address")
		ctx := r.Context()

		address, err := hc.resolveAddress(ctx, input)

		if err != nil {
			RenderError(w, http.StatusBadRequest, "invalid address")
			return
		}

		ctx = context.WithValue(ctx, common.ContextKeyAddress, address)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
