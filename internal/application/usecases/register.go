package usecases

import (
	"github.com/google/uuid"
	"github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/internal/domain/entities"
	ports2 "github.com/samuelorlato/football-api/internal/domain/ports/repositories"
	"github.com/samuelorlato/football-api/pkg/errs"
)

type registerUsecase struct {
	userRepository    ports2.UserRepository
	encryptionService ports.EncryptionService
}

func NewRegisterUsecase(userRepository ports2.UserRepository, encryptionService ports.EncryptionService) ports.RegisterUsecase {
	return &registerUsecase{
		userRepository,
		encryptionService,
	}
}

func (r *registerUsecase) Execute(registerRequest entities.RegisterRequest) error {
	user, err := r.userRepository.FindByUsername(registerRequest.Name)
	if err != nil {
		return errs.NewInternalServerError()
	}
	if user != nil {
		return errs.NewUnprocessableContentError("usuário já cadastrado")
	}

	userID := uuid.NewString()

	hashedPassword, err := r.encryptionService.HashPassword(registerRequest.Password)
	if err != nil {
		return errs.NewInternalServerError()
	}

	err = r.userRepository.Save(registerRequest.ToUserEntity(registerRequest, userID, *hashedPassword))
	if err != nil {
		return errs.NewInternalServerError()
	}

	return nil
}
