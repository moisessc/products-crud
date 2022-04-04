package main

import (
	"log"

	"products-crud/cmd/products-crud-api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
