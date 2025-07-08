package databases

import (
	"fmt"

	"github.com/samuelorlato/football-api/internal/infra/properties"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		properties.Properties().Database.Host,
		properties.Properties().Database.User,
		properties.Properties().Database.Password,
		properties.Properties().Database.Name,
		properties.Properties().Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
