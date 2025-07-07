package ports

import "time"

type TokenService interface {
	GenerateToken(userID string, expiresAt *time.Time, secret string) (*string, error)
	ValidateToken(token string, secret string) error
}
