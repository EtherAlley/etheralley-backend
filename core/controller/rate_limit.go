package controller

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/core/usecases"
)

func (hc *controller) rateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err := hc.verifyRateLimit.Do(ctx, &usecases.VerifyRateLimitInput{
			IpAddress: r.RemoteAddr,
		})

		if err != nil {
			hc.presenter.PresentTooManyRequests(w, r, err)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
