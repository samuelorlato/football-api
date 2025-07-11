package ports

import "github.com/samuelorlato/football-api/internal/domain/entities"

type FootballAPI interface {
	GetMatches(leagueCode string) ([]entities.Match, error)
	GetMatchdayMatches(leagueCode string, matchday int) ([]entities.Match, error)
	GetLeagues() ([]entities.League, error)
}
