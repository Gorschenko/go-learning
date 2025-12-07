package main

import (
	"fmt"
	"net/http"
	"test/configs"
	"test/internal/auth"
	"test/internal/links"
	"test/internal/stats"
	"test/internal/users"
	"test/packages/db"
	"test/packages/middlewares"
)

// 13.05 last lesson

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	usersRepository := users.NewUsersRepository(db)
	linksRepository := links.NewLinksRepository(db)
	statsRepository := stats.NewStatsRepository(db)

	// Services
	authService := auth.NewAuthService(usersRepository)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	links.NewLinksHandler(router, links.LinksHandlerDeps{
		Config:          conf,
		LinksRepository: linksRepository,
		StatsRepository: statsRepository,
	})

	// Middlewares
	stack := middlewares.ChainMiddlewares(
		middlewares.CorsMiddleware,
		middlewares.LoggerMiddleware,
	)

	server := http.Server{
		Addr:    ":" + conf.Server.Port,
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
