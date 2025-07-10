package ports

import "github.com/samuelorlato/football-api/internal/domain/entities"

type GetLeaguesUsecase interface {
	Execute() ([]entities.League, error)
}
