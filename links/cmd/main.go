package main

import (
	"fmt"
	"net/http"
	"test/configs"
	"test/internal/auth"
	links "test/internal/link"
	"test/packages/db"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()

	// handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	links.NewLinksHandler(router, links.LinksHandlerDeps{})

	server := http.Server{
		Addr:    ":" + conf.Server.Port,
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
