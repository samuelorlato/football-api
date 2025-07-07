package usecases

import (
	"os"
	"time"

	"github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/internal/domain/entities"
	ports2 "github.com/samuelorlato/football-api/internal/domain/ports/repositories"
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
		// TODO: create custom error
		return nil, err
	}

	err = l.encryptionService.CompareHashAndPassword(user.Password, password)
	if err != nil {
		// TODO: create custom error
		return nil, err
	}

	tokenExpirationTime := time.Now().Add(time.Hour * 24)
	// TODO: create secret
	tokenSecret := os.Getenv("JWT_SECRET")
	tokenString, err := l.tokenService.GenerateToken(user.ID, &tokenExpirationTime, tokenSecret)
	if err != nil {
		// TODO: create custom error
		return nil, err
	}

	token := &entities.Token{
		Token: *tokenString,
	}

	return token, nil
}
