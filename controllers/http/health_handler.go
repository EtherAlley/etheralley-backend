package http

import (
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (handler *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	RenderNoBody(w, http.StatusOK)
}
