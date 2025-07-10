package ports

import "github.com/samuelorlato/football-api/internal/domain/entities"

type FootballAPI interface {
	GetFinishedMatches(leagueCode string) ([]entities.Match, error)
	GetMatchdayFinishedMatches(leagueCode string, matchday int) ([]entities.Match, error)
	GetLeagues() ([]entities.League, error)
}
