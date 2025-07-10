package ports

import "github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"

type FootballController interface {
	GetMatches(leagueCode string, team *string, matchday *int) (*dtos.Matches, error)
	GetLeagues() ([]dtos.League, error)
}
