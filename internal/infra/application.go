package application

import (
	"errors"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/samuelorlato/football-api/internal/infra/dependency"
	"github.com/samuelorlato/football-api/internal/infra/properties"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	injector := dependency.Injector()

	err = injector.App.Start(":" + properties.Properties().Application.Port)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
