package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/controllers"
	httpControllers "github.com/etheralley/etheralley-core-api/controllers/http"
	"github.com/etheralley/etheralley-core-api/gateways/alchemy"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum"
	"github.com/etheralley/etheralley-core-api/gateways/mongo"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
	"github.com/etheralley/etheralley-core-api/gateways/thegraph"
	httpPresenters "github.com/etheralley/etheralley-core-api/presenters/http"
	"github.com/etheralley/etheralley-core-api/usecases"
	"go.uber.org/dig"
)

func awaitSigterm(logger common.ILogger) {
	ctx := context.Background()
	logger.Info(ctx, "awaiting sigterm")

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-cancelChan

	logger.Infof(ctx, "caught sigterm %v", sig)
}

func main() {
	container := dig.New()
	container.Provide(common.NewSettings)
	container.Provide(common.NewLogger)
	container.Provide(common.NewHttpClient)
	container.Provide(common.NewGraphQLClient)
	container.Provide(redis.NewGateway)
	container.Provide(mongo.NewGateway)
	container.Provide(ethereum.NewGateway)
	container.Provide(thegraph.NewGateway)
	container.Provide(alchemy.NewGateway)
	container.Provide(usecases.NewGetChallenge)
	container.Provide(usecases.NewGetProfile)
	container.Provide(usecases.NewGetDefaultProfile)
	container.Provide(usecases.NewSaveProfile)
	container.Provide(usecases.NewVerifyChallenge)
	container.Provide(usecases.NewGetNonFungibleToken)
	container.Provide(usecases.NewGetAllNonFungibleTokens)
	container.Provide(usecases.NewResolveENSAddress)
	container.Provide(usecases.NewResolveENSName)
	container.Provide(usecases.NewGetFungibleToken)
	container.Provide(usecases.NewGetAllFungibleTokens)
	container.Provide(usecases.NewGetStatistic)
	container.Provide(usecases.NewGetAllStatistics)
	container.Provide(usecases.NewGetInteractionUseCase)
	container.Provide(usecases.NewGetAllInteractionsUseCase)
	container.Provide(usecases.NewRecordProfileViewUseCase)
	container.Provide(usecases.NewGetTopProfilesUseCase)
	container.Provide(usecases.NewGetListingMetadata)
	container.Provide(usecases.NewGetListings)
	container.Provide(usecases.NewRefreshProfileUseCase)
	container.Provide(httpPresenters.NewPresenter)
	container.Provide(httpControllers.NewHttpController)

	// seed the random number generator based on app start time
	rand.Seed(time.Now().UnixNano())

	// start controllers in concurrent go routines
	err := container.Invoke(func(controller *httpControllers.HttpController) {
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
