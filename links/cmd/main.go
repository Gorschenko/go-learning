package main

import (
	"context"
	"fmt"
	"net/http"
	"test/configs"
	"test/internal/auth"
	"test/internal/links"
	"test/internal/users"
	"test/packages/db"
	"test/packages/middlewares"
	"time"
)

func tickOperation(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
		case <-ctx.Done():
			fmt.Println("Cancel")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go tickOperation(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

func main2() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	usersRepository := users.NewUsersRepository(db)
	linksRepository := links.NewLinksRepository(db)

	// Services
	authService := auth.NewAuthService(usersRepository)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	links.NewLinksHandler(router, links.LinksHandlerDeps{
		LinksRepository: linksRepository,
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
