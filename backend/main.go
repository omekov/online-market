package main

import (
	"log"
	"net/http"

	"github.com/omekov/online-market/backend/db"
	"github.com/omekov/online-market/backend/handlers"
)

func main() {
	data, err := db.Connection()
	if err != nil {
		log.Fatalf("Connection error: ", err)
	}
	defer data.Close()
	log.Println("Starting server port :5053")
	log.Fatal(http.ListenAndServe(":5053", handlers.Router()))
}
