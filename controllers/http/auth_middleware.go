package http

import (
	"net/http"
	"strings"

	"github.com/eflem00/go-example-app/usecases"
	"github.com/go-chi/chi/v5"
)

type AuthenticationMiddleware struct {
	authUseCase *usecases.AuthenticationUseCase
}

func NewAuthenticationMiddleware(authUseCase *usecases.AuthenticationUseCase) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		authUseCase,
	}
}

func (m *AuthenticationMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		address := chi.URLParam(r, "address")
		token := strings.Split(r.Header.Get("Authorization"), " ")

		if len(token) != 2 || token[0] != "Bearer" {
			RenderNoBody(w, http.StatusUnauthorized)
			return
		}

		challenge, err := m.authUseCase.GetChallenge(r.Context(), address)

		if err != nil {
			RenderNoBody(w, http.StatusUnauthorized)
			return
		}

		authentic := m.authUseCase.VerifySignature(address, token[1], challenge.Bytes())

		if !authentic {
			RenderNoBody(w, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})

}
