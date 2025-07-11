package dtos

import "github.com/samuelorlato/football-api/internal/domain/entities"

type RegisterFanResponse struct {
	ID    string `json:"id"`
	Name  string `json:"nome"`
	Email string `json:"email"`
	RegisterFanRequest
	Message string `json:"mensagem"`
}

func NewRegisterFanResponse(fan entities.Fan) *RegisterFanResponse {
	return &RegisterFanResponse{
		ID:    fan.ID,
		Name:  fan.Name,
		Email: fan.Email,
		RegisterFanRequest: RegisterFanRequest{
			Team: fan.Team,
		},
		Message: "Cadastro realizado com sucesso",
	}
}
