package controller

import (
	"context"
	"net/http"
	"time"
)

// cancel context after a certain threshold
//
// we are intentionally not returning an error response
// and allowing for the usecases to handle downstream context timeouts naturally
//
// see https://github.com/go-chi/chi/blob/master/middleware/timeout.go#L33
func (hc *controller) timeout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*30)

		defer cancel()

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
