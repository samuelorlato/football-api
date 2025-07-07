package ports

import "github.com/samuelorlato/football-api/internal/domain/entities"

type UserRepository interface {
	FindByUsername(username string) (*entities.User, error)
	Save(user entities.User) error
}
