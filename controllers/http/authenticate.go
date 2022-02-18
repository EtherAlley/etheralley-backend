package http

import (
	"net/http"
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/usecases"
)

func (hc *HttpController) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		address := ctx.Value(common.ContextKeyAddress).(string)
		token := strings.Split(r.Header.Get("Authorization"), " ")

		if len(token) != 2 || token[0] != "Bearer" {
			hc.presenter.PresentUnathorized(ctx, w)
			return
		}

		err := hc.verifyChallenge(r.Context(), &usecases.VerifyChallengeInput{
			Address: address,
			SigHex:  token[1],
		})

		if err != nil {
			hc.presenter.PresentUnathorized(ctx, w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
