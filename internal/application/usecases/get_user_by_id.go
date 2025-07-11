package usecases

import (
	"github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/internal/domain/entities"
	ports2 "github.com/samuelorlato/football-api/internal/domain/ports/repositories"
	"github.com/samuelorlato/football-api/pkg/errs"
)

type getUserByIDUsecase struct {
	userRepository ports2.UserRepository
}

func NewGetUserByIDUsecase(userRepository ports2.UserRepository) ports.GetUserByIDUsecase {
	return &getUserByIDUsecase{
		userRepository,
	}
}

func (g *getUserByIDUsecase) Execute(ID string) (*entities.User, error) {
	user, err := g.userRepository.FindByID(ID)
	if err != nil {
		return nil, errs.NewInternalServerError()
	}
	if user == nil {
		return nil, errs.NewNotFoundError("usuário não encontrado")
	}

	return user, nil
}
