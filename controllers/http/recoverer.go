package http

import (
	"net/http"
)

// Recoverer is a middleware that recovers from panics, logs the panic,
// and returns a HTTP 400 status to obfuscate internal errors from bad actors.
func (hc *HttpController) recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {

				hc.logger.Errorf(r.Context(), "Caught panic in recoverer: %+v", rvr)

				w.WriteHeader(http.StatusBadRequest)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
