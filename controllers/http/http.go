package http

import (
	"encoding/json"
	"net/http"

	"github.com/eflem00/go-example-app/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type HttpController struct {
	settings            *common.Settings
	logger              *common.Logger
	profileHandler      *ProfileHandler
	healthHandler       *HealthHandler
	recovererMiddleware *RecovererMiddleware
	challengeHandler    *ChallengeHandler
	authMiddleware      *AuthenticationMiddleware
}

func NewHttpController(settings *common.Settings, logger *common.Logger, profileHandler *ProfileHandler, healthHandler *HealthHandler, recovererMiddleware *RecovererMiddleware, challengeHandler *ChallengeHandler, authMiddleware *AuthenticationMiddleware) *HttpController {
	return &HttpController{
		settings,
		logger,
		profileHandler,
		healthHandler,
		recovererMiddleware,
		challengeHandler,
		authMiddleware,
	}
}

func (controller *HttpController) Start() error {
	controller.logger.Info("Starting http controller")

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://etheralley.io", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(controller.recovererMiddleware.Recoverer)

	r.Get("/", controller.healthHandler.Health)
	r.Get("/health", controller.healthHandler.Health)
	r.Get("/profiles/{address}", controller.profileHandler.GetProfileByAddress)
	r.With(controller.authMiddleware.Authenticate).Put("/profiles/{address}", controller.profileHandler.SaveProfile)
	r.Get("/challenge/{address}", controller.challengeHandler.GetChallenge)

	port := controller.settings.Port

	controller.logger.Infof("listening on %v", port)

	err := http.ListenAndServe(port, r)

	controller.logger.Err(err, "error in http controller")

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
