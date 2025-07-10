package dtos

import "github.com/samuelorlato/football-api/internal/domain/entities"

type MatchesResponse struct {
	Matches []Match `json:"matches"`
}

type Match struct {
	HomeTeam Team  `json:"homeTeam"`
	AwayTeam Team  `json:"awayTeam"`
	Score    Score `json:"score"`
}

type Team struct {
	ShortName string `json:"shortName"`
}

type Score struct {
	FullTime Time `json:"fullTime"`
}

type Time struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

func (m *MatchesResponse) ToEntities() []entities.Match {
	matches := make([]entities.Match, len(m.Matches))
	for i, match := range m.Matches {
		matches[i] = entities.Match{
			HomeTeam:  match.HomeTeam.ShortName,
			AwayTeam:  match.AwayTeam.ShortName,
			HomeScore: match.Score.FullTime.Home,
			AwayScore: match.Score.FullTime.Away,
		}
	}

	return matches
}
