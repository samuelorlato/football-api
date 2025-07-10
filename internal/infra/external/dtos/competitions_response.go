package dtos

import (
	"github.com/samuelorlato/football-api/internal/domain/entities"
)

type CompetitionsResponse struct {
	Competitions []Competition `json:"competitions"`
}

type Competition struct {
	Code          string        `json:"code"`
	Name          string        `json:"name"`
	CurrentSeason CurrentSeason `json:"currentSeason"`
}

type CurrentSeason struct {
	StartDate Date `json:"startDate"`
}

func (c *CompetitionsResponse) ToEntities() []entities.League {
	leagues := make([]entities.League, len(c.Competitions))
	for i, competition := range c.Competitions {
		leagues[i] = entities.League{
			ID:     competition.Code,
			Name:   competition.Name,
			Season: competition.CurrentSeason.StartDate.Year(),
		}
	}

	return leagues
}
