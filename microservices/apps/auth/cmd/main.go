package main

import (
	"log"
	"net/http"
	"strconv"

	"pkg/config"
	databases "pkg/databases/pg"
)

func main() {
	config, err := config.LoadConfig("../../config.json")
	if err != nil {
		panic(err)
	}

	_, err = databases.NewDb(config)

	if err != nil {
		panic(err)
	}

	log.Println("Starting Service 1 on :8080")
	address := config.Services.Auth.Host + ":" + strconv.Itoa(config.Services.Auth.Port)
	log.Fatal(http.ListenAndServe(address, nil))
}
