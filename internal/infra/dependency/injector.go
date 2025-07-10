package dependency

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samuelorlato/football-api/internal/application/services"
	"github.com/samuelorlato/football-api/internal/application/usecases"
	"github.com/samuelorlato/football-api/internal/infra/databases"
	"github.com/samuelorlato/football-api/internal/infra/external"
	"github.com/samuelorlato/football-api/internal/infra/properties"
	"github.com/samuelorlato/football-api/internal/infra/repositories"
	"github.com/samuelorlato/football-api/internal/infra/server/router"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/controllers"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/validators"
	"github.com/samuelorlato/football-api/internal/integration/persistance/models"
)

type injector struct {
	App *echo.Echo
}

func Injector() *injector {
	db := databases.NewPostgresConnection()
	db.AutoMigrate(&models.User{})
	userRepository := repositories.NewGormUserRepository(db)

	encryptionService := services.NewBcryptService()
	tokenService := services.NewJWTService()

	registerUsecase := usecases.NewRegisterUsecase(userRepository, encryptionService)
	loginUsecase := usecases.NewLoginUsecase(userRepository, encryptionService, tokenService)

	footballAPI := external.NewFootballAPI(&http.Client{}, properties.Properties().FootballAPI.BaseURL, properties.Properties().FootballAPI.Token)
	getLeaguesUsecase := usecases.NewGetLeaguesUsecase(footballAPI)
	getMatchesUsecase := usecases.NewGetMatchesUsecase(footballAPI)

	authorizationController := controllers.NewAuthorizationController(registerUsecase, loginUsecase)
	footballController := controllers.NewFootballController(getLeaguesUsecase, getMatchesUsecase)

	v10Validator := validators.NewV10Validator()

	app := router.New(footballController, authorizationController, v10Validator).Route()

	return &injector{
		App: app,
	}
}
