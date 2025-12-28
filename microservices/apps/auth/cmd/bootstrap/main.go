package main

import (
	"auth/internal/auth"
	"auth/internal/users"
	"log/slog"
	"net"
	"net/http"
	"strconv"

	"pkg/configs"
	"pkg/database"
	"pkg/logger"
	"pkg/middlewares"
)

func main() {
	config, err := configs.LoadConfig("../../config.json")
	if err != nil {
		panic(err)
	}

	logger.SetupLogger(&logger.LoggerServiceDependencies{
		Config: config,
	})

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
		Config:          config,
		UsersRepository: usersRepository,
	})

	// handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDependencies{
		AuthService: authService,
	})

	port := ":" + strconv.Itoa(config.Services.Auth.Port)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	listener.Close()

	handler := middlewares.CorrelationIdMiddleware(middlewares.LogsMiddleware(router))
	server := http.Server{
		Addr:    port,
		Handler: handler,
	}

	slog.Info(
		"Starting service",
		"Port",
		port,
	)
	server.ListenAndServe()
}
