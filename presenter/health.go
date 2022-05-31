package presenter

import (
	"net/http"
)

func (p *presenter) PresentHealth(w http.ResponseWriter, r *http.Request) {
	p.presentText(w, r, http.StatusOK, "OK")
}
