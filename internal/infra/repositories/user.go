package repositories

import (
	"github.com/samuelorlato/football-api/internal/domain/entities"
	ports "github.com/samuelorlato/football-api/internal/domain/ports/repositories"
)

type userRepository struct{}

func NewUserRepository() ports.UserRepository {
	return &userRepository{}
}

func (u *userRepository) FindByUsername(username string) (*entities.User, error) {
	return nil, nil
}

func (u *userRepository) Save(user entities.User) error {
	return nil
}
