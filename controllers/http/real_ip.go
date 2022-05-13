package http

import (
	"net/http"
	"strings"
)

var trueClientIP = http.CanonicalHeaderKey("True-Client-IP")
var xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
var xRealIP = http.CanonicalHeaderKey("X-Real-IP")

const localhost = "127.0.0.1"

// See https://github.com/go-chi/chi/blob/master/middleware/realip.go
func (hc *HttpController) realIP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tcip := r.Header.Get(trueClientIP)
		xRealIP := r.Header.Get(xRealIP)
		xForwardedFor := r.Header.Get(xForwardedFor)
		hc.logger.Info(r.Context()).Msgf("tcip: %v", tcip)
		hc.logger.Info(r.Context()).Msgf("realip: %v", xRealIP)
		hc.logger.Info(r.Context()).Msgf("xff: %v", xForwardedFor)

		// RemoteAddr is not in a recognizable format on dev without the 3 headers above
		// This makes things like recording profile views a little more pleasant on local while keeping validation in place
		if hc.settings.IsDev() {
			r.RemoteAddr = localhost
		} else if rip := getIP(r); rip != "" {
			r.RemoteAddr = rip
		}

		hc.logger.Info(r.Context()).Msgf("remoteaddr: %v", r.RemoteAddr)

		h.ServeHTTP(w, r)
	})
}

func getIP(r *http.Request) string {
	if xrip := r.Header.Get(xRealIP); xrip != "" {
		return xrip
	} else if xff := r.Header.Get(xForwardedFor); xff != "" {
		i := strings.Index(xff, ", ")
		if i == -1 {
			i = len(xff)
		}
		return xff[:i]
	}
	return ""
}
