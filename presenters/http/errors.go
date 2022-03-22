package http

import (
	"net/http"
)

func (p *httpPresenter) PresentBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	p.logger.Err(r.Context(), err, "bad request err")
	p.presentJSON(w, r, http.StatusBadRequest, toErrJson("bad request"))
}

func (p *httpPresenter) PresentUnathorized(w http.ResponseWriter, r *http.Request) {
	p.logger.Error(r.Context(), "unauthorized err")
	p.presentJSON(w, r, http.StatusUnauthorized, toErrJson("unathorized"))
}

func toErrJson(msg string) *errJson {
	return &errJson{
		Message: msg,
	}
}

type errJson struct {
	Message string `json:"message"`
}
