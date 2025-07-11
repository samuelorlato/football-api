package ports

import "github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"

type UserController interface {
	GetByID(ID string) (*dtos.User, error)
}
