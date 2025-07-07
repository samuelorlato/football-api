package ports

import "github.com/samuelorlato/football-api/internal/domain/entities"

type LoginUsecase interface {
	Execute(username, password string) (*entities.Token, error)
}
