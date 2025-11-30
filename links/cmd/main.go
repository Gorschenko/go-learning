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

	// repositories
	linksRepository := links.NewLinksRepository(db)

	// handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	links.NewLinksHandler(router, links.LinksHandlerDeps{
		LinksRepository: linksRepository,
	})

	server := http.Server{
		Addr:    ":" + conf.Server.Port,
		Handler: middlewares.CorsMiddleware(middlewares.LoggerMiddleware(router)),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
