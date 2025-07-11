package ports

import "github.com/samuelorlato/football-api/internal/domain/entities"

type UserRepository interface {
	FindByID(ID string) (*entities.User, error)
	FindByUsername(username string) (*entities.User, error)
	Save(user entities.User) error
}
