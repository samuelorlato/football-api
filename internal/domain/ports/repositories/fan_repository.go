package ports

import "github.com/samuelorlato/football-api/internal/domain/entities"

type FanRepository interface {
	FindByEmail(email string) (*entities.Fan, error)
	FindByTeam(team string) ([]entities.Fan, error)
	Save(fan entities.Fan) error
}
