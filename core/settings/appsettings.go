package settings

import (
	"os"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/joho/godotenv"
)

func NewAppSettings() common.IAppSettings {
	env := os.Getenv("ENV")

	if env == "dev" {
		err := godotenv.Load(".env/.env.core.dev")
		if err != nil {
			panic(err)
		}
	}

	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}

	return &appSettings{
		env:      env,
		appname:  "core-api",
		hostname: hostname,
		port:     os.Getenv("PORT"),
	}
}

type appSettings struct {
	env      string
	port     string
	appname  string
	hostname string
}

func (s *appSettings) Port() string {
	return s.port
}

func (s *appSettings) Appname() string {
	return s.appname
}

func (s *appSettings) Hostname() string {
	return s.hostname
}

func (s *appSettings) Env() string {
	return s.env
}

func (s *appSettings) IsDev() bool {
	return s.env == "dev"
}
