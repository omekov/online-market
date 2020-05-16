package sqlstore_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = dbConfigTest()
	os.Exit(m.Run())
}

// dbConfig - изъятие переменных окружений
func dbConfigTest() string {
	if _, err := os.Stat("test.env"); !os.IsNotExist(err) {
		err = godotenv.Load("test.env")
		if err != nil {
			log.Fatal(err)
		}
	}
	host, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		log.Fatal("env POSTGRES_HOST not found")
	}
	port, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		log.Fatal("env POSTGRES_PORT not found")
	}
	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		log.Fatal("env POSTGRES_USER not found")
	}
	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		log.Fatal("env POSTGRES_PASSWORD not found")
	}
	name, ok := os.LookupEnv("POSTGRES_NAME")
	if !ok {
		log.Fatal("env POSTGRES_NAME not found")
	}
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		name,
	)
}
