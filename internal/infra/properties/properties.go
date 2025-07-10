package properties

import "os"

type properties struct {
	Application *application
	Database    *database
	FootballAPI *footballAPI
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

type footballAPI struct {
	BaseURL string
	Token   string
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
		FootballAPI: &footballAPI{
			BaseURL: os.Getenv("FOOTBALL_API_BASE_URL"),
			Token:   os.Getenv("FOOTBALL_API_TOKEN"),
		},
	}
}
