package http

import (
	"net/http"
)

func (p *httpPresenter) PresentHealth(w http.ResponseWriter, r *http.Request) {
	p.presentText(w, r, http.StatusOK, "OK")
}
