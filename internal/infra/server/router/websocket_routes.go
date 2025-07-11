package router

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/samuelorlato/football-api/internal/infra/properties"
	"github.com/samuelorlato/football-api/internal/infra/server/middlewares"
	"github.com/samuelorlato/football-api/internal/infra/server/websocket"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
)

func (r *router) handleFan(c echo.Context) error {
	conn, err := websocket.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userID := claims["user_id"].(string)
	email := claims["email"].(string)
	if userID == "" || email == "" {
		conn.WriteJSON("token inválido")
		websocket.RemoveConnection(userID)
		conn.Close()
	}

	_, err = r.userController.GetByID(userID)
	if err != nil {
		conn.WriteJSON("o token não pertence mais a um usuário")
		websocket.RemoveConnection(userID)
		conn.Close()
	}

	fan, err := r.fanController.GetByEmail(email)
	if err != nil {
		conn.WriteJSON(err.Error())
		websocket.RemoveConnection(userID)
		conn.Close()
	}

	websocket.RegisterConnection(userID, fan.Team, conn)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			websocket.RemoveConnection(userID)
			conn.Close()
			break
		}
	}

	return nil
}

func (r *router) handleBroadcast(c echo.Context) error {
	conn, err := websocket.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userID := claims["user_id"].(string)
	email := claims["email"].(string)
	if userID == "" || email == "" {
		conn.WriteJSON("token inválido")
		websocket.RemoveConnection(userID)
		conn.Close()
	}

	_, err = r.userController.GetByID(userID)
	if err != nil {
		conn.WriteJSON("o token não pertence mais a um usuário")
		websocket.RemoveConnection(userID)
		conn.Close()
	}

	for {
		var payload dtos.BroadcastRequest
		err := conn.ReadJSON(&payload)
		if err != nil {
			conn.Close()
			break
		}

		err = r.validator.Struct(payload)
		if err != nil {
			conn.WriteJSON(r.validator.GetErrors(err, &payload))
			continue
		}

		r.broadcastController.Broadcast(payload, websocket.GetConnectionsByTeam(payload.Team))
	}

	return nil
}

func (r *router) mountWebSocketRoutes(e *echo.Echo) {
	ws := e.Group("/ws")
	ws.Use(echojwt.JWT([]byte(properties.Properties().Application.JWTSecret)))
	ws.GET("/torcedor", r.handleFan)
	ws.GET("/admin/broadcast", r.handleBroadcast, middlewares.RequireRole("admin"))
}
