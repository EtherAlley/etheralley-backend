package http

import (
	"encoding/json"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type HttpController struct {
	settings               *common.Settings
	logger                 *common.Logger
	getProfileUsecase      usecases.IGetProfileUsecase
	saveProfileUsecase     usecases.ISaveProfileUseCase
	getChallengeUsecase    usecases.IGetChallengeUseCase
	verifyChallengeUseCase usecases.IVerifyChallengeUseCase
}

func NewHttpController(settings *common.Settings, logger *common.Logger, getProfileUsecase *usecases.GetProfileUsecase, saveProfileUsecase *usecases.SaveProfileUseCase, getChallengeUsecase *usecases.GetChallengeUseCase, verifyChallengeUseCase *usecases.VerifyChallengeUseCase) *HttpController {
	return &HttpController{
		settings,
		logger,
		getProfileUsecase,
		saveProfileUsecase,
		getChallengeUsecase,
		verifyChallengeUseCase,
	}
}

func (hc *HttpController) Start() error {
	hc.logger.Info("Starting http controller")

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
	r.Route("/profiles", hc.registerProfileRoutes)
	r.Route("/challenge", hc.registerChallengeRoutes)

	port := hc.settings.Port

	hc.logger.Infof("listening on %v", port)

	err := http.ListenAndServe(port, r)

	hc.logger.Err(err, "error in http controller")

	return err
}

func (controller *HttpController) Exit() {
	controller.logger.Error("detected exit in http controller")
}

type ErrBody struct {
	Message string `json:"message"`
}

func RenderErr(w http.ResponseWriter, statusCode int, msg string) {
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
