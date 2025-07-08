package dependency

import (
	"github.com/labstack/echo/v4"
	"github.com/samuelorlato/football-api/internal/application/services"
	"github.com/samuelorlato/football-api/internal/application/usecases"
	"github.com/samuelorlato/football-api/internal/infra/repositories"
	"github.com/samuelorlato/football-api/internal/infra/server/router"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/controllers"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/validators"
)

type injector struct {
	App *echo.Echo
}

func Injector() *injector {
	userRepository := repositories.NewUserRepository()
	encryptionService := services.NewBcryptService()
	registerUsecase := usecases.NewRegisterUsecase(userRepository, encryptionService)
	tokenService := services.NewJWTService()
	loginUsecase := usecases.NewLoginUsecase(userRepository, encryptionService, tokenService)
	authorizationController := controllers.NewAuthorizationController(registerUsecase, loginUsecase)
	v10Validator := validators.NewV10Validator()

	app := router.New(authorizationController, v10Validator).Route()

	return &injector{
		App: app,
	}
}
