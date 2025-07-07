package router

import (
	"github.com/labstack/echo/v4"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
	"github.com/samuelorlato/football-api/internal/integration/ports"
)

type router struct {
	authorizationController ports.AuthorizationController
}

func New(authorizationController ports.AuthorizationController) *router {
	return &router{
		authorizationController,
	}
}

func (r *router) Route() *echo.Echo {
	e := echo.New()

	a := e.Group("/auth")
	a.POST("/cadastro", func(c echo.Context) error {
		var registerRequest dtos.RegisterRequest
		err := c.Bind(&registerRequest)
		if err != nil {
			return err
		}

		err = r.authorizationController.Register(registerRequest)
		if err != nil {
			return err
		}

		return c.JSON(204, nil)
	})
	a.POST("/login", func(c echo.Context) error {
		var loginRequest dtos.LoginRequest
		err := c.Bind(&loginRequest)
		if err != nil {
			return err
		}

		token, err := r.authorizationController.Login(loginRequest)
		if err != nil {
			return err
		}

		return c.JSON(200, token)
	})

	return e
}
