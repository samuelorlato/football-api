package router

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/samuelorlato/football-api/internal/infra/properties"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
)

func (r *router) handleFans(c echo.Context) error {
	var registerFanRequest dtos.RegisterFanRequest
	err := c.Bind(&registerFanRequest)
	if err != nil {
		return badRequest("corpo da requisição inválido")
	}

	err = r.validator.Struct(registerFanRequest)
	if err != nil {
		return validationErrors(r.validator, err, &registerFanRequest)
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userID := claims["user_id"].(string)
	name := claims["name"].(string)
	email := claims["email"].(string)
	if userID == "" || name == "" || email == "" {
		return badRequest("token inválido")
	}

	_, err = r.userController.GetByID(userID)
	if err != nil {
		return badRequest("o token não pertence mais a um usuário")
	}

	fan, err := r.broadcastController.Subscribe(registerFanRequest, name, email)
	if err != nil {
		return handleAppError(err)
	}

	return c.JSON(http.StatusCreated, fan)
}

func (r *router) mountFansRoutes(e *echo.Echo) {
	t := e.Group("/torcedores")
	t.Use(echojwt.JWT([]byte(properties.Properties().Application.JWTSecret)))
	t.POST("", r.handleFans)
}
