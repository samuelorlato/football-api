package ports

import "github.com/samuelorlato/football-api/internal/domain/entities"

type SubscribeUsecase interface {
	Execute(registerFanRequest entities.RegisterFanRequest) (*entities.Fan, error)
}
