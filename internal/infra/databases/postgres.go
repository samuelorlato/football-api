package databases

import (
	"fmt"
	"time"

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

	var db *gorm.DB
	var err error

	for i := range 10 {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Printf("retry %d to connect to database", i+1)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		panic(err)
	}

	return db
}
