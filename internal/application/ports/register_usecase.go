package ports

import "github.com/samuelorlato/football-api/internal/domain/entities"

type RegisterUsecase interface {
	Execute(registerRequest entities.RegisterRequest) error
}
