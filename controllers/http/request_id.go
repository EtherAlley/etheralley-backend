package http

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/etheralley/etheralley-core-api/common"
)

var RequestIDHeader = "X-Request-Id"
var reqid uint64

// add a request id to the context
//
// if one is not present in headers, generate it.
//
// see https://github.com/go-chi/chi/blob/master/middleware/request_id.go#L67
func (hc *HttpController) requestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			requestID = fmt.Sprintf("%06d", atomic.AddUint64(&reqid, 1))
		}
		ctx = context.WithValue(ctx, common.ContextKeyRequestId, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
