package dtos

type LoginRequest struct {
	User     string `json:"usuário" validate:"required"`
	Password string `json:"senha" validate:"required"`
}
