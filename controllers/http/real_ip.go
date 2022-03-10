package http

import (
	"net/http"
	"strings"
)

var trueClientIP = http.CanonicalHeaderKey("True-Client-IP")
var xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
var xRealIP = http.CanonicalHeaderKey("X-Real-IP")

const localhost = "127.0.0.1"

// TODO: Ideally these headers can be trusted from the load balancer of the cloud provider we settle on
//
// See https://github.com/go-chi/chi/blob/master/middleware/realip.go
func (hc *HttpController) realIP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// RemoteAddr is not in a recognizable format on dev without the 3 headers above
		// This makes things like recording profile views a little more pleasant on local while keep validation in place
		if hc.settings.IsDev() {
			r.RemoteAddr = localhost
		} else if rip := getIP(r); rip != "" {
			r.RemoteAddr = rip
		}
		h.ServeHTTP(w, r)
	})
}

func getIP(r *http.Request) string {
	if tcip := r.Header.Get(trueClientIP); tcip != "" {
		return tcip
	} else if xrip := r.Header.Get(xRealIP); xrip != "" {
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
