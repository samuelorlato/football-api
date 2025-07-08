package properties

import "os"

type properties struct {
	Application *application
	Database    *database
}

type application struct {
	Port      string
	JWTSecret string
}

type database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Properties() *properties {
	return &properties{
		Application: &application{
			Port:      os.Getenv("APP_PORT"),
			JWTSecret: os.Getenv("JWT_SECRET"),
		},
		Database: &database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
}
