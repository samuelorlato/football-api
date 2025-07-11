package ports

import "github.com/samuelorlato/football-api/internal/domain/entities"

type GetUserByIDUsecase interface {
	Execute(id string) (*entities.User, error)
}
