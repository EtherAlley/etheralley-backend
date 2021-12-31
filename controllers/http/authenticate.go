package http

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		address := chi.URLParam(r, "address")
		token := strings.Split(r.Header.Get("Authorization"), " ")

		if len(token) != 2 || token[0] != "Bearer" {
			RenderError(w, http.StatusBadRequest, "unauthorized")
			return
		}

		err := hc.verifyChallengeUseCase(r.Context(), address, token[1])

		if err != nil {
			RenderError(w, http.StatusBadRequest, "unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
