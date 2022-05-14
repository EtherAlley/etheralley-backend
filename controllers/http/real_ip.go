package http

import (
	"net"
	"net/http"
	"strings"
)

var xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")

// See https://github.com/go-chi/chi/blob/master/middleware/realip.go
func (hc *HttpController) realIP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hc.logger.Info(r.Context()).Msgf("x-forwarded-for: %v", xForwardedFor)
		hc.logger.Info(r.Context()).Msgf("remoteaddr before: %v", r.RemoteAddr)

		if ip := getIP(r); ip != "" {
			r.RemoteAddr = ip
		}

		if host, err := parseIP(r); err == nil {
			r.RemoteAddr = host
		}

		hc.logger.Info(r.Context()).Msgf("remoteaddr after: %v", r.RemoteAddr)

		h.ServeHTTP(w, r)
	})
}

// if xff header present, pick last addr in comma delimited list
func getIP(r *http.Request) (ip string) {
	xff := r.Header.Get(xForwardedFor)

	if xff == "" {
		return ""
	}

	i := strings.LastIndex(xff, ", ")

	if i+2 >= len(xff) {
		return ""
	}

	if i == -1 {
		return xff
	}

	return xff[i+2:]
}

// validate the ip format
// parse the host out of the ip addr
func parseIP(r *http.Request) (string, error) {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	return host, err
}
