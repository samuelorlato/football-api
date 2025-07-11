package router

import (
	"github.com/labstack/echo/v4"
	"github.com/samuelorlato/football-api/internal/integration/ports"
)

type router struct {
	userController          ports.UserController
	fanController           ports.FanController
	broadcastController     ports.BroadcastController
	footballController      ports.FootballController
	authorizationController ports.AuthorizationController
	validator               ports.Validator
}

func New(userController ports.UserController, fanController ports.FanController, broadcastController ports.BroadcastController, footballController ports.FootballController, authorizationController ports.AuthorizationController, validator ports.Validator) *router {
	return &router{
		userController,
		fanController,
		broadcastController,
		footballController,
		authorizationController,
		validator,
	}
}

func (r *router) Route() *echo.Echo {
	e := echo.New()

	r.mountAuthRoutes(e)
	r.mountLeagueRoutes(e)
	r.mountFansRoutes(e)
	r.mountWebSocketRoutes(e)

	return e
}
