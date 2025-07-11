package models

import (
	"time"

	"github.com/samuelorlato/football-api/internal/domain/entities"
)

type Fan struct {
	ID        string `gorm:"primaryKey;type:uuid"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Team      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (f *Fan) ToEntity() entities.Fan {
	return entities.Fan{
		ID:    f.ID,
		Name:  f.Name,
		Email: f.Email,
		Team:  f.Team,
	}
}

func (f *Fan) FromEntity(fan entities.Fan) {
	f.ID = fan.ID
	f.Name = fan.Name
	f.Email = fan.Email
	f.Team = fan.Team
}
