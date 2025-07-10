package dtos

import "github.com/samuelorlato/football-api/internal/domain/entities"

type Matches struct {
	Matchday int     `json:"rodada"`
	Matches  []Match `json:"partidas"`
}

type Match struct {
	HomeTeam string `json:"time_casa"`
	AwayTeam string `json:"time_fora"`
	Score    string `json:"placar"`
}

func NewMatch(match entities.Match, score string) Match {
	return Match{
		HomeTeam: match.HomeTeam,
		AwayTeam: match.AwayTeam,
		Score:    score,
	}
}
