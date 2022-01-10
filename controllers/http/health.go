package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) registerHealthRoutes(r chi.Router) {
	r.Get("/", hc.healthRoute)
}

func (hc *HttpController) healthRoute(w http.ResponseWriter, r *http.Request) {
	RenderNoBody(w, http.StatusOK)
}
