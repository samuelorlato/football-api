package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func RequireRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			roleInToken, ok := claims["role"].(string)
			if !ok || roleInToken != role {
				return echo.NewHTTPError(http.StatusForbidden, "acesso negado: permissão insuficiente")
			}

			return next(c)
		}
	}
}
