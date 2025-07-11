package dtos

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	*jwt.RegisteredClaims
}
