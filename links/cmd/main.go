package main

import (
	"fmt"
	"net/http"
	"test/configs"
	"test/internal/auth"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	server:= http.Server{
		Addr: ":" + conf.Server.Port,
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}   