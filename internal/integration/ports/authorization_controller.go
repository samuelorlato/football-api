package ports

import "github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"

type AuthorizationController interface {
	Register(dtos.RegisterRequest) error
	Login(dtos.LoginRequest) (*dtos.Token, error)
}
