package dtos

type LoginRequest struct {
	User     string `json:"usuario" validate:"required"`
	Password string `json:"senha" validate:"required"`
}
