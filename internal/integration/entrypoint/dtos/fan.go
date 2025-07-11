package dtos

import "github.com/samuelorlato/football-api/internal/domain/entities"

type Fan struct {
	ID    string
	Name  string
	Email string
	Team  string
}

func NewFan(fan *entities.Fan) *Fan {
	return &Fan{
		ID:    fan.ID,
		Name:  fan.Name,
		Email: fan.Email,
		Team:  fan.Team,
	}
}
