package db

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/lib/pq" //postgres driver
)

// DB ...
var DB *sql.DB

// Connection ...
func Connection() {
	set := newSettings()
	db, err := sql.Open("postgres",
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
		log.Fatalf("SQL Open: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("SQL Ping: %v", err)
	}
	DB = db
}
