package presenter

import (
	"net/http"
)

func (p *presenter) PresentBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	p.logger.Info(r.Context()).Err(err).Msg("bad request")
	p.presentJSON(w, r, http.StatusBadRequest, toErrJson("bad request"))
}

func (p *presenter) PresentUnathorized(w http.ResponseWriter, r *http.Request, err error) {
	p.logger.Info(r.Context()).Err(err).Msg("unauthorized")
	p.presentJSON(w, r, http.StatusUnauthorized, toErrJson("unathorized"))
}

func (p *presenter) PresentNotFound(w http.ResponseWriter, r *http.Request, err error) {
	p.logger.Info(r.Context()).Err(err).Msg("not found")
	p.presentJSON(w, r, http.StatusNotFound, toErrJson("not found"))
}

func (p *presenter) PresentTooManyRequests(w http.ResponseWriter, r *http.Request, err error) {
	p.logger.Warn(r.Context()).Err(err).Msg("too many requests")
	p.presentJSON(w, r, http.StatusTooManyRequests, toErrJson("too many requests"))
}

func (p *presenter) PresentForbiddenRequest(w http.ResponseWriter, r *http.Request, err error) {
	p.logger.Info(r.Context()).Err(err).Msg("forbidden request")
	p.presentJSON(w, r, http.StatusForbidden, toErrJson("forbidden request"))
}

func toErrJson(msg string) *errJson {
	return &errJson{
		Message: msg,
	}
}

type errJson struct {
	Message string `json:"message"`
}
