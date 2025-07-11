package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/pkg/errs"
)

type JWTService struct{}

func NewJWTService() ports.TokenService {
	return &JWTService{}
}

type claim struct {
	userID string
	*jwt.RegisteredClaims
}

func (j *JWTService) GenerateToken(userID string, expiresAt *time.Time, secret string) (*string, error) {
	claims := claim{
		userID:           userID,
		RegisteredClaims: &jwt.RegisteredClaims{},
	}

	if expiresAt != nil {
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(*expiresAt)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (j *JWTService) ValidateToken(tokenString string, secret string) error {
	claims := &claim{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return errs.NewUnauthorizedError("token inválido ou expirado")
	}

	return nil
}
