package app

import (
	"devices/internal/devices"
	"net/http"
	"pkg/configs"
	"pkg/database"
	"pkg/logger"
	pkg_mqtt "pkg/mqtt"
)

func GetApp(configPath string) (http.Handler, *configs.Config) {
	config, err := configs.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	db, err := database.NewDb(config)
	if err != nil {
		panic(err)
	}

	logger.SetupLogger(config)
	router := http.NewServeMux()

	mqttService, err := pkg_mqtt.NewMqttService(config)
	if err != nil {
		panic(err)
	}

	// repositories
	devices.NewDevicesRepository(&database.RepositoryDependencies{
		Database: db,
		Config:   config,
	})

	// services

	// handlers
	devices.NewMqttDevicesHandler(mqttService)

	return router, config
}
