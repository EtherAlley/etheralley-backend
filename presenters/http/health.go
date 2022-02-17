package http

import (
	"context"
	"net/http"
)

func (p *httpPresenter) PresentHealth(ctx context.Context, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("."))
}
