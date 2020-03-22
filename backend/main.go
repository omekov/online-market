package main

import (
	"log"
	"net/http"

	"github.com/omekov/online-market/backend/db"
	"github.com/omekov/online-market/backend/handlers"
)

func init() {
	db.Connection()
}
func main() {
	defer db.DB.Close()
	log.Fatal(http.ListenAndServe(":5053", handlers.Router()))
}
