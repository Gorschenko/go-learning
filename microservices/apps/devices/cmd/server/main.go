package main

import (
	"devices/internal/app"
	"log/slog"
	"net"
	"net/http"
	"strconv"
)

func main() {
	app, config := app.GetApp("../../config.json")

	port := ":" + strconv.Itoa(config.Services.Devices.Port)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	listener.Close()

	server := http.Server{
		Addr:    port,
		Handler: app,
	}

	slog.Info(
		"Starting service",
		"Port",
		port,
	)
	server.ListenAndServe()
}
