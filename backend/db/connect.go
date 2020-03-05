package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "marketdb"
)
func config() {

}
func Start() {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password,	dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("sql open", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("sql ping", err)
	}
	fmt.Println("Successfully connected DataBase!")
}

func Connection() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password,	dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}