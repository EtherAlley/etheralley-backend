package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (hc *HttpController) registerHealthRoutes(r chi.Router) {
	r.Get("/", hc.health)
}

func (hc *HttpController) health(w http.ResponseWriter, r *http.Request) {
	RenderNoBody(w, http.StatusOK)
}
