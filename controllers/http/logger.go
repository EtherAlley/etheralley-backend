package http

import (
	"net/http"
	"time"
)

// TODO: get response status code
// log details of the request/response
//
// see https://github.com/go-chi/chi/blob/master/middleware/logger.go
func (hc *HttpController) logEvent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()

		defer func() {
			hc.logger.Event(r.Context(), []struct {
				Key   string
				Value string
			}{
				{Key: "method", Value: r.Method},
				{Key: "path", Value: r.URL.Path},
				{Key: "resptime", Value: time.Since(t1).String()},
				{Key: "remoteaddr", Value: r.RemoteAddr},
				{Key: "hostname", Value: hc.settings.Hostname()},
				{Key: "instanceid", Value: hc.settings.InstanceID()},
			})
		}()

		next.ServeHTTP(w, r)
	})
}
