package models

import (
	"time"

	"github.com/samuelorlato/football-api/internal/domain/entities"
)

type User struct {
	ID           string `gorm:"primaryKey;type:uuid"`
	Name         string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) ToEntity() entities.User {
	return entities.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Role:         u.Role,
	}
}

func (u *User) FromEntity(user entities.User) {
	u.ID = user.ID
	u.Name = user.Name
	u.Email = user.Email
	u.PasswordHash = user.PasswordHash
	u.Role = user.Role
}
