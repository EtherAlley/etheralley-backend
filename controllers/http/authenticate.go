package http

import (
	"net/http"
	"strings"
)

func (hc *HttpController) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		address := r.Context().Value(contextKeyAddress).(string)
		token := strings.Split(r.Header.Get("Authorization"), " ")

		if len(token) != 2 || token[0] != "Bearer" {
			RenderError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		err := hc.verifyChallenge(r.Context(), address, token[1])

		if err != nil {
			RenderError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
