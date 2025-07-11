package controllers

import (
	"github.com/gorilla/websocket"
	ports2 "github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
	"github.com/samuelorlato/football-api/internal/integration/ports"
)

type broadcastController struct {
	subscribeUsecase ports2.SubscribeUsecase
}

func NewBroadcastController(subscribeUsecase ports2.SubscribeUsecase) ports.BroadcastController {
	return &broadcastController{
		subscribeUsecase,
	}
}

func (b *broadcastController) Subscribe(registerFanRequest dtos.RegisterFanRequest, name, email string) (*dtos.RegisterFanResponse, error) {
	fan, err := b.subscribeUsecase.Execute(registerFanRequest.ToEntity(name, email))
	if err != nil {
		return nil, err
	}

	return dtos.NewRegisterFanResponse(*fan), nil
}

func (b *broadcastController) Broadcast(payload dtos.BroadcastRequest, connections []*websocket.Conn) {
	for _, conn := range connections {
		conn.WriteJSON(payload)
	}
}
