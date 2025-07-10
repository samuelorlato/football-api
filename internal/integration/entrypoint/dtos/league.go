package dtos

import "github.com/samuelorlato/football-api/internal/domain/entities"

type League struct {
	ID     string `json:"id"`
	Name   string `json:"nome"`
	Season int    `json:"temporada"`
}

func NewLeague(league entities.League) League {
	return League{
		ID:     league.ID,
		Name:   league.Name,
		Season: league.Season,
	}
}
