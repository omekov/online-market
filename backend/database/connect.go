package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // ...
	_ "github.com/lib/pq"      // ...
)

var db *sql.DB

// Settings ...
type Settings struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// GetEnv ...
func GetEnv(ev string, defVal ...string) string {
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
func NewSettings() *Settings {
	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return &Settings{
		Host:     GetEnv("DBHOST"),
		Port:     GetEnv("DBPORT"),
		User:     GetEnv("DBUSER"),
		Password: GetEnv("DBPASS"),
		DBName:   GetEnv("DBNAME"),
	}
}

// Connection ...
func Connection() {
	set := NewSettings()
	data, err := sql.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			set.Host,
			set.Port,
			set.User,
			set.Password,
			set.DBName,
		),
	)
	if err != nil {
		log.Fatalf("sql open", err)
	}
	defer data.Close()
	err = data.Ping()
	if err != nil {
		log.Fatalf("sql ping", err)
	}
	log.Println("Successfully connected DataBase!")
	db = data
}
