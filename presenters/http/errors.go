package http

import (
	"context"
	"net/http"
)

func (p *httpPresenter) PresentBadRequest(ctx context.Context, w http.ResponseWriter, err error) {
	p.logger.Err(ctx, err, "bad request err")
	render(w, http.StatusBadRequest, toErrJson("bad request"))
}

func (p *httpPresenter) PresentUnathorized(ctx context.Context, w http.ResponseWriter) {
	p.logger.Error(ctx, "unauthorized err")
	render(w, http.StatusUnauthorized, toErrJson("unathorized"))
}

func toErrJson(msg string) *errJson {
	return &errJson{
		Message: msg,
	}
}

type errJson struct {
	Message string `json:"message"`
}
