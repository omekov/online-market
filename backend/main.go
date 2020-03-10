package main

import (
	"log"
	"net/http"

	"github.com/omekov/online-market/backend/database"
	"github.com/omekov/online-market/backend/handlers"
)

func init() {
	database.Connection()
}
func main() {
	log.Fatal(http.ListenAndServe(":5053", handlers.Router()))
}
