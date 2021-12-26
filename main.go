package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/controllers"
	"github.com/eflem00/go-example-app/controllers/http"
	"github.com/eflem00/go-example-app/gateways/cache"
	"github.com/eflem00/go-example-app/gateways/mongo"
	"github.com/eflem00/go-example-app/usecases"
	"go.uber.org/dig"
)

func awaitSigterm(logger *common.Logger) {
	logger.Info("awaiting sigterm")

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-cancelChan

	logger.Infof("caught sigterm %v", sig)
}

func main() {
	container := dig.New()
	container.Provide(common.NewSettings)
	container.Provide(common.NewLogger)
	container.Provide(cache.NewCache)
	container.Provide(mongo.NewDb)
	container.Provide(mongo.NewProfileRepository)
	container.Provide(usecases.NewProfileUseCase)
	container.Provide(http.NewRecovererMiddleware)
	container.Provide(http.NewHealthHandler)
	container.Provide(http.NewProfileHandler)
	container.Provide(http.NewHttpController)

	// start controllers in concurrent go routines
	err := container.Invoke(func(controller *http.HttpController) {
		go controllers.StartController(controller)
	})

	if err != nil {
		panic(err)
	}

	// blocking call in main routine to await sigterm
	err = container.Invoke(awaitSigterm)

	if err != nil {
		panic(err)
	}

	// TODO: Shutdown gracefully below
}
