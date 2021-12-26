package http

import (
	"net/http"

	"github.com/eflem00/go-example-app/common"
)

type RecovererMiddleware struct {
	logger *common.Logger
}

func NewRecovererMiddleware(logger *common.Logger) *RecovererMiddleware {
	return &RecovererMiddleware{
		logger,
	}
}

// Recoverer is a middleware that recovers from panics, logs the panic,
// and returns a HTTP 400 status to obfuscate internal errors from bad actors.
func (m *RecovererMiddleware) Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {

				m.logger.Errorf("Caught panic in recoverer: %+v", rvr)

				w.WriteHeader(http.StatusBadRequest)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
