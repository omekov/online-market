package db

import (
	"log"
	"os"

	"github.com/joho/godotenv" // ...
)

// Settings ...
type settings struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// GetEnv ...
func getEnv(ev string, defVal ...string) string {
	v := os.Getenv(ev)
	if v == "" {
		if len(defVal) == 0 {
			log.Fatalf("Not exists require env variable %s", ev)
		}
		v = defVal[0]
	}
	return v
}

// NewSettings ...
func newSettings() *settings {
	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return &settings{
		Host:     getEnv("DBHOST"),
		Port:     getEnv("DBPORT"),
		User:     getEnv("DBUSER"),
		Password: getEnv("DBPASS"),
		DBName:   getEnv("DBNAME"),
	}
}
