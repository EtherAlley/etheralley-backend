package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/daos/controller"
	"github.com/etheralley/etheralley-backend/daos/settings"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()
	container.Provide(context.Background)
	container.Provide(settings.NewAppSettings)
	container.Provide(common.NewLogger)
	container.Provide(controller.NewHttpController)

	// seed the random number generator based on app start time
	rand.Seed(time.Now().UnixNano())

	// start the http controller inside a go routine
	if err := container.Invoke(startHttpController); err != nil {
		panic(err)
	}

	// blocking call in main go routine to await sigterm
	if err := container.Invoke(awaitSigterm); err != nil {
		panic(err)
	}
}

func startHttpController(ctx context.Context, controller controller.IHttpController) {
	go func() {
		err := controller.Start(ctx)
		panic(err)
	}()
}

func awaitSigterm(ctx context.Context, logger common.ILogger) {
	logger.Info(ctx).Msg("awaiting sigterm")

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-cancelChan

	logger.Info(ctx).Msgf("caught sigterm %v", sig)
}
