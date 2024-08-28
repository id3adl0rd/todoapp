package config

import (
	"github.com/joho/godotenv"
	"os"
	"to-do-app/logger"
)

var (
	Config *Params
)

type Params struct {
	Postgres
}

type Postgres struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

func NewConfig() *Params {
	if err := godotenv.Load(); err != nil {
		logger.Log.Errorf("failed to load env variables: %v", err)
		return nil
	}

	return &Params{
		Postgres: Postgres{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			SSLMode:  os.Getenv("SSL_MODE"),
			TimeZone: os.Getenv("TIMEZONE"),
		},
	}
}
