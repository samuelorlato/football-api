package ports

import "github.com/samuelorlato/football-api/internal/domain/entities"

type GetMatchesUseCase interface {
	Execute(leagueCode string, team *string, matchday *int) ([]entities.Match, error)
}
