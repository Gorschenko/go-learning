package main

import (
	"fmt"
	"log"
	"net/http"

	"shared/models"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := models.User{ID: 1, Name: "From Service 1"}
		fmt.Fprintf(w, "Service 1 Response\nUser: %v", user)
	})

	log.Println("Starting Service 1 on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
