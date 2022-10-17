package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/go-chi/chi/v5"
)

type IHttpController interface {
	Start(context.Context) error
}

type httpController struct {
	logger      common.ILogger
	appSettings common.IAppSettings
}

func NewHttpController(logger common.ILogger, appSettings common.IAppSettings) IHttpController {
	return &httpController{
		logger,
		appSettings,
	}
}

func (hc *httpController) Start(ctx context.Context) error {
	hc.logger.Info(ctx).Msg("starting http controller")

	r := chi.NewRouter()

	r.Get("/", hc.healthRoute)

	port := hc.appSettings.Port()

	hc.logger.Info(ctx).Msgf("listening on port %v", port)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), r)

	hc.logger.Error(ctx).Err(err).Msg("error in http controller")

	return err
}

func (hc *httpController) healthRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
