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
	eventsbus "test/packages/events"
	"test/packages/middlewares"
)

// Пропущены 3 и 4 модули
// 14.05 последняя лекция

func App() http.Handler {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := eventsbus.NewEventBus()

	// Repositories
	usersRepository := users.NewUsersRepository(db)
	linksRepository := links.NewLinksRepository(db)
	statsRepository := stats.NewStatsRepository(db)

	// Services
	authService := auth.NewAuthService(usersRepository)
	statsService := stats.NewStatsService(&stats.StatsServiceDeps{
		EventBus:        eventBus,
		StatsRepository: statsRepository,
	})

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	links.NewLinksHandler(router, links.LinksHandlerDeps{
		Config:          conf,
		LinksRepository: linksRepository,
		EventBus:        eventBus,
	})
	stats.NewStatsHandler(router, stats.StatsHandlerDeps{
		Config:          conf,
		StatsRepository: statsRepository,
	})

	go statsService.AddClick()

	// Middlewares
	stack := middlewares.ChainMiddlewares(
		middlewares.CorsMiddleware,
		middlewares.LoggerMiddleware,
	)

	return stack(router)
}

func main() {
	app := App()

	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
