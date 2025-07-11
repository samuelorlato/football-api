package dtos

import "github.com/samuelorlato/football-api/internal/domain/entities"

type RegisterRequest struct {
	Name     string `json:"usuário" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"senha" validate:"required"`
	Role     string `json:"papel" validate:"required,oneof=admin torcedor"`
}

func (r *RegisterRequest) ToEntity() entities.RegisterRequest {
	return entities.RegisterRequest{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		Role:     r.Role,
	}
}
