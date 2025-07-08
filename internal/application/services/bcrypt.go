package services

import (
	"github.com/samuelorlato/football-api/internal/application/ports"
	"golang.org/x/crypto/bcrypt"
)

type bcryptService struct{}

func NewBcryptService() ports.EncryptionService {
	return &bcryptService{}
}

func (e *bcryptService) HashPassword(password string) (*string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	hashedPasswordString := string(hashedPassword)

	return &hashedPasswordString, nil
}

func (e *bcryptService) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
