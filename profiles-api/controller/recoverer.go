package controller

import (
	"fmt"
	"net/http"
)

// Recoverer is a middleware that recovers from panics, logs the panic,
// and returns a HTTP 400 status to obfuscate internal errors from bad actors.
func (hc *controller) recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {

				err := fmt.Errorf("caught panic in recoverer: %+v", rvr)

				hc.logger.Error(r.Context()).Err(err).Send()

				hc.presenter.PresentBadRequest(w, r, err)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
