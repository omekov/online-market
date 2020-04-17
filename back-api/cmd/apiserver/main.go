package main

import (
	"log"

	"github.com/omekov/online-market/back-api/internal/app/apiserver"
)

func main() {
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
