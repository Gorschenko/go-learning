package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Starting Service 2 on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
