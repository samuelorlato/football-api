package ports

import "github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"

type FanController interface {
	GetByEmail(email string) (*dtos.Fan, error)
}
