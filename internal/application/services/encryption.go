package services

import (
	"github.com/samuelorlato/football-api/internal/application/ports"
	"golang.org/x/crypto/bcrypt"
)

type encryptionService struct{}

func NewEncryptionService() ports.EncryptionService {
	return &encryptionService{}
}

func (e *encryptionService) HashPassword(password string) (*string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	hashedPasswordString := string(hashedPassword)

	return &hashedPasswordString, nil
}

func (e *encryptionService) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
