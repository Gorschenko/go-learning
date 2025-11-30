package main

import (
	"fmt"
	"net/http"
	"test/configs"
	"test/internal/auth"
	"test/internal/links"
	"test/packages/db"
	"test/packages/middlewares"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linksRepository := links.NewLinksRepository(db)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	links.NewLinksHandler(router, links.LinksHandlerDeps{
		LinksRepository: linksRepository,
	})

	// Middlewares
	stack := middlewares.ChainMiddlewares(
		middlewares.CorsMiddleware,
		middlewares.LoggerMiddleware,
		middlewares.IsAuthenticated,
	)

	server := http.Server{
		Addr:    ":" + conf.Server.Port,
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
