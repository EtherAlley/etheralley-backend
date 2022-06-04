package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/etheralley/etheralley-apis/common"
	"github.com/etheralley/etheralley-apis/core/controller"
	"github.com/etheralley/etheralley-apis/core/gateways"
	"github.com/etheralley/etheralley-apis/core/gateways/ethereum"
	"github.com/etheralley/etheralley-apis/core/gateways/mongo"
	"github.com/etheralley/etheralley-apis/core/gateways/offchain"
	"github.com/etheralley/etheralley-apis/core/gateways/redis"
	"github.com/etheralley/etheralley-apis/core/gateways/thegraph"
	"github.com/etheralley/etheralley-apis/core/presenter"
	"github.com/etheralley/etheralley-apis/core/settings"
	"github.com/etheralley/etheralley-apis/core/usecases"
	"go.uber.org/dig"
)

func main() {
	// build the dependency container
	container := dig.New()
	container.Provide(context.Background)
	container.Provide(settings.NewAppSettings)
	container.Provide(common.NewLogger)
	container.Provide(common.NewHttpClient)
	container.Provide(common.NewGraphQLClient)
	container.Provide(settings.NewSettings)
	container.Provide(redis.NewGateway)
	container.Provide(mongo.NewGateway)
	container.Provide(ethereum.NewGateway)
	container.Provide(thegraph.NewGateway)
	container.Provide(offchain.NewGateway)
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
	container.Provide(usecases.NewGetCurrency)
	container.Provide(usecases.NewGetAllCurrenciesUseCase)
	container.Provide(usecases.NewGetStoreMetadata)
	container.Provide(usecases.NewVerifyRateLimit)
	container.Provide(presenter.NewHttpPresenter)
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

func startHttpController(ctx context.Context, logger common.ILogger, controller controller.IHttpController, cacheGateway gateways.ICacheGateway, databaseGateway gateways.IDatabaseGateway, offchainGateway gateways.IOffchainGateway) {
	// do any further initialization here
	logger.Info(ctx).Msg("initializing gateways")

	if err := cacheGateway.Init(ctx); err != nil {
		panic(err)
	}

	if err := databaseGateway.Init(ctx); err != nil {
		panic(err)
	}

	if err := offchainGateway.Init(ctx); err != nil {
		panic(err)
	}

	go func() {
		// start is intended to be a blocking call
		// if start returns, one of our controllers is no longer active and thus we should force a panic
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
