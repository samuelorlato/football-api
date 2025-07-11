package ports

import "github.com/samuelorlato/football-api/internal/domain/entities"

type GetFanByEmailUsecase interface {
	Execute(email string) (*entities.Fan, error)
}
