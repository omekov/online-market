package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/omekov/online-market/backend/db"
	"github.com/omekov/online-market/backend/handlers"
)

func init() {
	fmt.Println(time.Now())
	db.Connection()
}
func main() {
	defer db.DB.Close()
	log.Fatal(http.ListenAndServe(":5053", handlers.Router()))
}
