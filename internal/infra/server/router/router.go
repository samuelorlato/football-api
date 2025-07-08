package router

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
	"github.com/samuelorlato/football-api/internal/integration/ports"
	"github.com/samuelorlato/football-api/pkg/errs"
)

type router struct {
	authorizationController ports.AuthorizationController
	validator               ports.Validator
}

func New(authorizationController ports.AuthorizationController, validator ports.Validator) *router {
	return &router{
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

	return e
}
