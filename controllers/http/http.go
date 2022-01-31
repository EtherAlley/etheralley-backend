package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type HttpController struct {
	settings            *common.Settings
	logger              *common.Logger
	getProfile          usecases.IGetProfileUseCase
	saveProfile         usecases.ISaveProfileUseCase
	getChallenge        usecases.IGetChallengeUseCase
	verifyChallenge     usecases.IVerifyChallengeUseCase
	getNonFungibleToken usecases.IGetNonFungibleTokenUseCase
	resolveAddress      usecases.IResolveAddressUseCase
	getFungibleToken    usecases.IGetFungibleTokenUseCase
	getStatistic        usecases.IGetStatisticUseCase
}

func NewHttpController(
	settings *common.Settings,
	logger *common.Logger,
	getProfile usecases.IGetProfileUseCase,
	saveProfile usecases.ISaveProfileUseCase,
	getChallenge usecases.IGetChallengeUseCase,
	verifyChallenge usecases.IVerifyChallengeUseCase,
	getNonFungibleToken usecases.IGetNonFungibleTokenUseCase,
	resolveAddress usecases.IResolveAddressUseCase,
	getFungibleToken usecases.IGetFungibleTokenUseCase,
	getStatistic usecases.IGetStatisticUseCase,
) *HttpController {
	return &HttpController{
		settings,
		logger,
		getProfile,
		saveProfile,
		getChallenge,
		verifyChallenge,
		getNonFungibleToken,
		resolveAddress,
		getFungibleToken,
		getStatistic,
	}
}

func (hc *HttpController) Start() error {
	hc.logger.Info("starting http controller")

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://etheralley.io", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(hc.recoverer)

	r.Route("/health", hc.registerHealthRoutes)

	r.Route("/profiles/{address}", func(r chi.Router) {
		r.Use(hc.resolveENSName)
		r.Get("/", hc.getProfileByAddressRoute)
		r.With(hc.authenticate).Put("/", hc.saveProfileRoute)
	})

	r.Route("/challenges/{address}", func(r chi.Router) {
		r.Use(hc.resolveENSName)
		r.Get("/", hc.getChallengeRoute)
	})

	r.Route("/", func(r chi.Router) {
		r.Use(hc.parseContract)
		r.Get("/token", hc.getTokenRoute)
		r.Get("/nft", hc.getNFTRoute)
		r.Get("/statistic", hc.getStatisticRoute)
	})

	port := hc.settings.Port

	hc.logger.Infof("listening on port %v", port)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), r)

	hc.logger.Err(err, "error in http controller")

	return err
}

func (controller *HttpController) Exit() {
	controller.logger.Error("detected exit in http controller")
}

// context keys
type contextKey string

func (c contextKey) String() string {
	return "etheralley context key " + string(c)
}

var (
	contextKeyAddress  = contextKey("address")
	contextKeyContract = contextKey("contract")
)

// response rendering
type ErrBody struct {
	Message string `json:"message"`
}

func RenderErr(w http.ResponseWriter, statusCode int, err error) {
	Render(w, statusCode, ErrBody{Message: err.Error()})
}

func RenderError(w http.ResponseWriter, statusCode int, msg string) {
	Render(w, statusCode, ErrBody{Message: msg})
}

func RenderNoBody(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

func Render(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}
