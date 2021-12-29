package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/controllers"
	"github.com/etheralley/etheralley-core-api/controllers/http"
	"github.com/etheralley/etheralley-core-api/gateways/mongo"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
	"github.com/etheralley/etheralley-core-api/usecases"
	"go.uber.org/dig"
)

func awaitSigterm(logger *common.Logger) {
	logger.Info("awaiting sigterm")

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-cancelChan

	logger.Infof("caught sigterm %v", sig)
}

func setRandSeed() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	container := dig.New()
	container.Provide(common.NewSettings)
	container.Provide(common.NewLogger)
	container.Provide(redis.NewGateway)
	container.Provide(mongo.NewGateway)
	container.Provide(usecases.NewGetChallengeUseCase)
	container.Provide(usecases.NewGetProfileUsecase)
	container.Provide(usecases.NewSaveProfileUseCase)
	container.Provide(usecases.NewVerifyChallengeUseCase)
	container.Provide(http.NewHttpController)

	setRandSeed()

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
