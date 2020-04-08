package main

import (
	"log"

	"github.com/omekov/online-market/go-app/internal/app/apiserver"
)

func main() {
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
