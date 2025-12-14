package main

import (
	"auth/internal/auth"
	"auth/internal/users"
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

	db, err := database.NewDb(config)
	if err != nil {
		panic(err)
	}

	router := http.NewServeMux()

	// repositories
	usersRepository := users.NewUsersRepository(&users.UsersRepositoryDependencies{
		Database: db,
		Config:   config,
	})

	// services
	authService := auth.NewAuthService(auth.AuthServiceDependencies{
		UsersRepository: usersRepository,
	})

	// controllers
	auth.NewAuthController(router, auth.AuthControllerDependencies{
		AuthService: authService,
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
