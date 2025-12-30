package app

import (
	"auth/internal/auth"
	"auth/internal/cars"
	"auth/internal/users"
	"net/http"
	"pkg/configs"
	"pkg/database"
	"pkg/logger"
	"pkg/middlewares"
	"time"
)

func GetApp(configPath string) (http.Handler, *configs.Config) {
	config, err := configs.LoadConfig(configPath)
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
	usersRepository := users.NewUsersRepository(&database.RepositoryDependencies{
		Database: db,
		Config:   config,
	})
	_ = cars.NewCarsRepository(&database.RepositoryDependencies{
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

	middlewaresStack := middlewares.CombainMiddlewares(
		middlewares.CorrelationIdMiddleware,
		middlewares.LogsMiddleware,
		middlewares.TimeoutMiddleware(5*time.Second),
	)

	return middlewaresStack(router), config
}
