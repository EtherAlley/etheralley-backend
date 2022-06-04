package controller

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/etheralley/etheralley-apis/common"
)

var RequestIDHeader = "X-Request-Id"

// add a request id to the context
//
// if one is not present in headers, generate it.
//
// see https://github.com/go-chi/chi/blob/master/middleware/request_id.go#L67
func (hc *controller) requestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			requestID = fmt.Sprintf("%10d", rand.Intn(10000000000))
		}
		ctx = context.WithValue(ctx, common.ContextKeyRequestId, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
