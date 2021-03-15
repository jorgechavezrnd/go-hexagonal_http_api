package main

import (
	"log"

	"github.com/jorgechavezrnd/go-hexagonal_http_api/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
