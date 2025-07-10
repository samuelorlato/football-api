package router

import (
	"errors"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/samuelorlato/football-api/internal/infra/properties"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
	"github.com/samuelorlato/football-api/internal/integration/ports"
	"github.com/samuelorlato/football-api/pkg/errs"
)

type router struct {
	footballController      ports.FootballController
	authorizationController ports.AuthorizationController
	validator               ports.Validator
}

func New(footballController ports.FootballController, authorizationController ports.AuthorizationController, validator ports.Validator) *router {
	return &router{
		footballController,
		authorizationController,
		validator,
	}
}

func (r *router) Route() *echo.Echo {
	e := echo.New()

	a := e.Group("/auth")
	a.POST("/cadastro", func(c echo.Context) error {
		var registerRequest dtos.RegisterRequest
		err := c.Bind(&registerRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"erro": "corpo da requisição inválido",
			})
		}

		err = r.validator.Struct(registerRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, r.validator.GetErrors(err, &registerRequest))
		}

		err = r.authorizationController.Register(registerRequest)
		if err != nil {
			var customErr *errs.Error
			if errors.As(err, &customErr) {
				return c.JSON(customErr.Code, echo.Map{
					"erro": customErr.Message,
				})
			}

			unexpectedErr := errs.NewInternalServerError()
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"erro": unexpectedErr.Message,
			})
		}

		return c.JSON(http.StatusCreated, nil)
	})
	a.POST("/login", func(c echo.Context) error {
		var loginRequest dtos.LoginRequest
		err := c.Bind(&loginRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"erro": "corpo da requisição inválido",
			})
		}

		err = r.validator.Struct(loginRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, r.validator.GetErrors(err, &loginRequest))
		}

		token, err := r.authorizationController.Login(loginRequest)
		if err != nil {
			var customErr *errs.Error
			if errors.As(err, &customErr) {
				return c.JSON(customErr.Code, echo.Map{
					"erro": customErr.Message,
				})
			}

			unexpectedErr := errs.NewInternalServerError()
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"erro": unexpectedErr.Message,
			})
		}

		return c.JSON(http.StatusOK, token)
	})

	c := e.Group("/campeonatos")
	c.Use(echojwt.JWT([]byte(properties.Properties().Application.JWTSecret)))
	c.GET("", func(c echo.Context) error {
		leagues, err := r.footballController.GetLeagues()
		if err != nil {
			var customErr *errs.Error
			if errors.As(err, &customErr) {
				return c.JSON(customErr.Code, echo.Map{
					"erro": customErr.Message,
				})
			}

			unexpectedErr := errs.NewInternalServerError()
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"erro": unexpectedErr.Message,
			})
		}

		return c.JSON(http.StatusOK, leagues)
	})
	c.GET("/:leagueCode/partidas", func(c echo.Context) error {
		var matchFilterRequest dtos.MatchFilterRequest
		err := c.Bind(&matchFilterRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"erro": "corpo da requisição inválido",
			})
		}

		err = r.validator.Struct(matchFilterRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, r.validator.GetErrors(err, &matchFilterRequest))
		}

		matches, err := r.footballController.GetMatches(matchFilterRequest.LeagueCode, matchFilterRequest.Team, matchFilterRequest.Matchday)
		if err != nil {
			var customErr *errs.Error
			if errors.As(err, &customErr) {
				return c.JSON(customErr.Code, echo.Map{
					"erro": customErr.Message,
				})
			}

			unexpectedErr := errs.NewInternalServerError()
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"erro": unexpectedErr.Message,
			})
		}

		return c.JSON(http.StatusOK, matches)
	})

	return e
}
