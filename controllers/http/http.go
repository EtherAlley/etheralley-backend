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
	getProfile          usecases.GetProfileUseCase
	saveProfile         usecases.SaveProfileUseCase
	getChallenge        usecases.GetChallengeUseCase
	verifyChallenge     usecases.VerifyChallengeUseCase
	getNonFungibleToken usecases.GetNonFungibleTokenUseCase
	getValidAddress     usecases.GetValidAddressUseCase
	getFungibleToken    usecases.GetFungibleTokenUseCase
}

func NewHttpController(
	settings *common.Settings,
	logger *common.Logger,
	getProfile usecases.GetProfileUseCase,
	saveProfile usecases.SaveProfileUseCase,
	getChallenge usecases.GetChallengeUseCase,
	verifyChallenge usecases.VerifyChallengeUseCase,
	getNonFungibleToken usecases.GetNonFungibleTokenUseCase,
	getValidAddress usecases.GetValidAddressUseCase,
	getFungibleToken usecases.GetFungibleTokenUseCase,
) *HttpController {
	return &HttpController{
		settings,
		logger,
		getProfile,
		saveProfile,
		getChallenge,
		verifyChallenge,
		getNonFungibleToken,
		getValidAddress,
		getFungibleToken,
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

	r.Route("/address/{address}", func(r chi.Router) {
		r.Use(hc.resolveENSName)
		r.Route("/profile", hc.registerProfileRoutes)
		r.Route("/challenge", hc.registerChallengeRoutes)
		r.Route("/token", hc.registerTokenRoutes)
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
	contextKeyAddress = contextKey("address")
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
