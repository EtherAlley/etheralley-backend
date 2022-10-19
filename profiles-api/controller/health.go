package controller

import (
	"net/http"
)

func (hc *controller) healthRoute(w http.ResponseWriter, r *http.Request) {
	hc.presenter.PresentHealth(w, r)
}
