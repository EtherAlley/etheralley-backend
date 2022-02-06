package http

import (
	"net/http"
)

func (hc *HttpController) healthRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("."))
}
