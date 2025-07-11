package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
)

func (r *router) handleRegister(c echo.Context) error {
	var registerRequest dtos.RegisterRequest
	err := c.Bind(&registerRequest)
	if err != nil {
		return badRequest("corpo da requisição inválido")
	}

	err = r.validator.Struct(registerRequest)
	if err != nil {
		return validationErrors(r.validator, err, &registerRequest)
	}

	err = r.authorizationController.Register(registerRequest)
	if err != nil {
		return handleAppError(err)
	}

	return c.NoContent(http.StatusCreated)
}

func (r *router) handleLogin(c echo.Context) error {
	var loginRequest dtos.LoginRequest
	err := c.Bind(&loginRequest)
	if err != nil {
		return badRequest("corpo da requisição inválido")
	}

	err = r.validator.Struct(loginRequest)
	if err != nil {
		return validationErrors(r.validator, err, &loginRequest)
	}

	token, err := r.authorizationController.Login(loginRequest)
	if err != nil {
		return handleAppError(err)
	}

	return c.JSON(http.StatusOK, token)
}

func (r *router) mountAuthRoutes(e *echo.Echo) {
	a := e.Group("/auth")
	a.POST("/cadastro", r.handleRegister)
	a.POST("/login", r.handleLogin)
}
