package dtos

import "github.com/samuelorlato/football-api/internal/domain/entities"

type User struct {
	ID string `json:"id"`
}

func NewUser(user *entities.User) *User {
	return &User{
		ID: user.ID,
	}
}
