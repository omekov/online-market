package main

import (
	"log"

	"github.com/omekov/online-market/backend/golang/internal/app/apiserver"
)

func main() {
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
