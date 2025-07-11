package ports

import (
	"github.com/gorilla/websocket"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
)

type BroadcastController interface {
	Subscribe(registerFanRequest dtos.RegisterFanRequest, name, email string) (*dtos.RegisterFanResponse, error)
	Broadcast(payload dtos.BroadcastRequest, connections []*websocket.Conn)
}
