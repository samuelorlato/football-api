package dtos

import (
	"time"

	"github.com/samuelorlato/football-api/internal/domain/entities"
)

type Matches struct {
	Matchday *int    `json:"rodada,omitempty"`
	Matches  []Match `json:"partidas"`
}

type Match struct {
	UTCDate  time.Time `json:"data"`
	HomeTeam string    `json:"time_casa"`
	AwayTeam string    `json:"time_fora"`
	Score    *string   `json:"placar,omitempty"`
}

func NewMatch(match entities.Match, score *string) Match {
	loc, _ := time.LoadLocation("America/Sao_Paulo")

	return Match{
		UTCDate:  match.UTCDate.In(loc),
		HomeTeam: match.HomeTeam,
		AwayTeam: match.AwayTeam,
		Score:    score,
	}
}
