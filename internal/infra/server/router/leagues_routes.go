package router

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/samuelorlato/football-api/internal/infra/properties"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
)

func (r *router) handleLeagues(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userID := claims["user_id"].(string)
	if userID == "" {
		return badRequest("token inválido")
	}

	_, err := r.userController.GetByID(userID)
	if err != nil {
		return badRequest("o token não pertence mais a um usuário")
	}

	leagues, err := r.footballController.GetLeagues()
	if err != nil {
		return handleAppError(err)
	}

	return c.JSON(http.StatusOK, leagues)
}

func (r *router) handleLeagueMatches(c echo.Context) error {
	var matchFilterRequest dtos.MatchFilterRequest
	err := c.Bind(&matchFilterRequest)
	if err != nil {
		return badRequest("corpo da requisição inválido")
	}

	err = r.validator.Struct(matchFilterRequest)
	if err != nil {
		return validationErrors(r.validator, err, &matchFilterRequest)
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userID := claims["user_id"].(string)
	if userID == "" {
		return badRequest("token inválido")
	}

	_, err = r.userController.GetByID(userID)
	if err != nil {
		return badRequest("o token não pertence mais a um usuário")
	}

	matches, err := r.footballController.GetMatches(matchFilterRequest.LeagueCode, matchFilterRequest.Team, matchFilterRequest.Matchday)
	if err != nil {
		return handleAppError(err)
	}

	return c.JSON(http.StatusOK, matches)
}

func (r *router) mountLeagueRoutes(e *echo.Echo) {
	c := e.Group("/campeonatos")
	c.Use(echojwt.JWT([]byte(properties.Properties().Application.JWTSecret)))
	c.GET("", r.handleLeagues)
	c.GET("/:leagueCode/partidas", r.handleLeagueMatches)
}
