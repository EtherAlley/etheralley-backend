package http

import (
	"net/http"
)

func (hc *HttpController) healthRoute(w http.ResponseWriter, r *http.Request) {
	hc.presenter.PresentHealth(w, r)
}
