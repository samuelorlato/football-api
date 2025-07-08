package usecases

import (
	"time"

	"github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/internal/domain/entities"
	ports2 "github.com/samuelorlato/football-api/internal/domain/ports/repositories"
	"github.com/samuelorlato/football-api/internal/infra/properties"
	"github.com/samuelorlato/football-api/pkg/errs"
)

type loginUsecase struct {
	userRepository    ports2.UserRepository
	encryptionService ports.EncryptionService
	tokenService      ports.TokenService
}

func NewLoginUsecase(userRepository ports2.UserRepository, encryptionService ports.EncryptionService, tokenService ports.TokenService) ports.LoginUsecase {
	return &loginUsecase{
		userRepository,
		encryptionService,
		tokenService,
	}
}

func (l *loginUsecase) Execute(username, password string) (*entities.Token, error) {
	user, err := l.userRepository.FindByUsername(username)
	if err != nil {
		return nil, errs.NewInternalServerError()
	}
	if user == nil {
		return nil, errs.NewNotFoundError("usuário não encontrado")
	}

	err = l.encryptionService.CompareHashAndPassword(user.Password, password)
	if err != nil {
		return nil, errs.NewBadRequestError("senha incorreta")
	}

	tokenExpirationTime := time.Now().Add(time.Hour * 24)
	tokenSecret := properties.Properties().Application.JWTSecret
	tokenString, err := l.tokenService.GenerateToken(user.ID, &tokenExpirationTime, tokenSecret)
	if err != nil {
		return nil, errs.NewInternalServerError()
	}

	token := &entities.Token{
		Token: *tokenString,
	}

	return token, nil
}
