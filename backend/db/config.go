package db

import (
	"os"

	"log"

	"github.com/joho/godotenv" // ...
)

// settings ...
type settings struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// NewSettings ...
func newSettings() *settings {
	env := os.Getenv("PROD_ENV")
	if env == "local" {
		if _, err := os.Stat(".env.local"); !os.IsNotExist(err) {
			err = godotenv.Load(".env.local")
			if err != nil {
				log.Fatal("Error loading .env.local file")
			}
		}
	} else {
		if _, err := os.Stat(".env"); !os.IsNotExist(err) {
			err = godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}
		}
	}
	return &settings{
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASS"),
		DBName:   os.Getenv("DBNAME"),
	}
}
