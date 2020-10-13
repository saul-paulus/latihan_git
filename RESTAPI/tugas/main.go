package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
