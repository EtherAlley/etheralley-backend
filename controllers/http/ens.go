package http

import (
	"context"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) resolveENSName(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		address := chi.URLParam(r, "address")
		ctx := r.Context()

		address, err := hc.resolveAddress(ctx, address)

		if err != nil {
			RenderError(w, http.StatusBadRequest, "invalid address")
			return
		}

		ctx = context.WithValue(ctx, common.ContextKeyAddress, address)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
