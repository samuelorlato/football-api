package dtos

import "github.com/samuelorlato/football-api/internal/domain/entities"

type RegisterFanRequest struct {
	Team string `json:"time" validate:"required"`
}

func (r *RegisterFanRequest) ToEntity(name, email string) entities.RegisterFanRequest {
	return entities.RegisterFanRequest{
		Name:  name,
		Email: email,
		Team:  r.Team,
	}
}
