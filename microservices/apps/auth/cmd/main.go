package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"

	"pkg/configs"
	"pkg/database"
)

func main() {
	config, err := configs.LoadConfig("../../config.json")
	if err != nil {
		panic(err)
	}

	_, err = database.NewDb(config)
	if err != nil {
		panic(err)
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Hi")
	})

	address := ":" + strconv.Itoa(config.Services.Auth.Port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	listener.Close()

	server := http.Server{
		Addr:    address,
		Handler: router,
	}

	log.Printf("Starting Auth service on %s port", address)
	server.ListenAndServe()

}
