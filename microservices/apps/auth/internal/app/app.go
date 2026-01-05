package app

import (
	"auth/internal/auth"
	"auth/internal/users"
	"net/http"
	"pkg/cache"
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

	logger.SetupLogger(config)

	db, err := database.NewDb(config)
	if err != nil {
		panic(err)
	}

	router := http.NewServeMux()

	cacheRepository, err := cache.NewCacheRepository(config)

	if err != nil {
		panic(err)
	}

	// cache repositories
	cacheUsersRepository := cache.NewCacheUsersRepository(&cache.CacheUsersRepositoryDependencies{
		CacheRepository: cacheRepository,
	})

	// db repositories
	usersRepository := users.NewUsersRepository(&database.RepositoryDependencies{
		Database: db,
		Config:   config,
	})

	// services
	authService := auth.NewAuthService(&auth.AuthServiceDependencies{
		Config:          config,
		UsersRepository: usersRepository,
	})
	usersService := users.NewUsersService(&users.UsersServiceDependencies{
		UsersRepository:      usersRepository,
		CacheUsersRepository: cacheUsersRepository,
	})

	// handlers
	auth.NewAuthHandler(router, &auth.AuthHandlerDependencies{
		AuthService: authService,
	})
	users.NewUsersHandler(router, &users.UsersHandlerDependencies{
		UsersService: usersService,
	})

	middlewaresStack := middlewares.CombainMiddlewares(
		middlewares.CorrelationIdMiddleware,
		middlewares.LogsMiddleware,
		middlewares.TimeoutMiddleware(5*time.Second),
	)

	return middlewaresStack(router), config
}
