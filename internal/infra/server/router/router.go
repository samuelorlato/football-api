package router

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/samuelorlato/football-api/internal/infra/properties"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
	"github.com/samuelorlato/football-api/internal/integration/ports"
	"github.com/samuelorlato/football-api/pkg/errs"
)

type router struct {
	fanController           ports.FanController
	broadcastController     ports.BroadcastController
	footballController      ports.FootballController
	authorizationController ports.AuthorizationController
	validator               ports.Validator
}

func New(fanController ports.FanController, broadcastController ports.BroadcastController, footballController ports.FootballController, authorizationController ports.AuthorizationController, validator ports.Validator) *router {
	return &router{
		fanController,
		broadcastController,
		footballController,
		authorizationController,
		validator,
	}
}

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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var userConnections = make(map[string]*websocket.Conn)
var userSubscriptions = make(map[string]string)

func registerConnection(user string, team string, conn *websocket.Conn) {
	userConnections[user] = conn
	userSubscriptions[user] = team
}

func removeConnection(user string) {
	delete(userConnections, user)
	delete(userSubscriptions, user)
}

func getConnectionsByTeam(team string) []*websocket.Conn {
	var connections []*websocket.Conn
	for userID, subscribedTeam := range userSubscriptions {
		if subscribedTeam == team {
			conn := userConnections[userID]
			connections = append(connections, conn)
		}
	}
	return connections
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

	t := e.Group("/torcedores")
	t.Use(echojwt.JWT([]byte(properties.Properties().Application.JWTSecret)))
	t.POST("", func(c echo.Context) error {
		var registerFanRequest dtos.RegisterFanRequest
		err := c.Bind(&registerFanRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"erro": "corpo da requisição inválido",
			})
		}

		err = r.validator.Struct(registerFanRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, r.validator.GetErrors(err, &registerFanRequest))
		}

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		fan, err := r.broadcastController.Subscribe(registerFanRequest, claims["name"].(string), claims["email"].(string))
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

		return c.JSON(http.StatusCreated, fan)
	})

	ws := e.Group("/ws")
	ws.Use(echojwt.JWT([]byte(properties.Properties().Application.JWTSecret)))
	ws.GET("/torcedor", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string)
		email := claims["email"].(string)
		fan, err := r.fanController.GetByEmail(email)
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

		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}

		registerConnection(userID, fan.Team, conn)

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				removeConnection(userID)
				conn.Close()
				break
			}
		}

		return nil
	})
	ws.GET("/admin/broadcast", func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
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

			r.broadcastController.Broadcast(payload, getConnectionsByTeam(payload.Team))
		}

		return nil
	}, RequireRole("admin"))

	return e
}
