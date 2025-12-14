package main

import (
	"fmt"
	"log"
	"net/http"

	"pkg/models"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := models.User{ID: 2, Name: "From Service 2"}
		fmt.Fprintf(w, "Service 2 Response\nUser: %v", user)
	})

	log.Println("Starting Service 2 on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
