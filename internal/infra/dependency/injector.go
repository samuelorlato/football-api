package dependency

import (
	"github.com/labstack/echo/v4"
	"github.com/samuelorlato/football-api/internal/infra/server/router"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/controllers"
	"github.com/samuelorlato/football-api/internal/integration/ports"
)

type injector struct {
	authorizationController ports.AuthorizationController
	App                     *echo.Echo
}

func Injector() *injector {
	authorizationController := controllers.AuthorizationController()

	app := router.New(authorizationController).Route()

	return &injector{
		authorizationController: authorizationController,
		App:                     app,
	}
}
