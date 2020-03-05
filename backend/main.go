package main

import (
	"log"
	"net/http"

	"github.com/omekov/online-market/backend/handlers"

	"github.com/omekov/online-market/backend/db"
)

func main() {
	db.Start()
	log.Println("Starting server port :5500")
	log.Fatal(http.ListenAndServe(":5500", handlers.Router()))
}
