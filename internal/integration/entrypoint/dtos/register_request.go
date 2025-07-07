package dtos

import "github.com/samuelorlato/football-api/internal/domain/entities"

type RegisterRequest struct {
	Name     string `json:"usuario"`
	Email    string `json:"email"`
	Password string `json:"senha"`
}

func (r *RegisterRequest) ToEntity() entities.RegisterRequest {
	return entities.RegisterRequest{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}
