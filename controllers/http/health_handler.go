package http

import (
	"net/http"

	"github.com/go-chi/render"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (handler *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
}
