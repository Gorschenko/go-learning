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

	logger.SetupLogger(config)

	_, err = database.NewDb(config)
	if err != nil {
		panic(err)
	}

	mqttService, err := pkg_mqtt.NewMqttService(config)
	if err != nil {
		panic(err)
	}

	// handlers
	devices.NewMqttDevicesHandler(mqttService)

	router := http.NewServeMux()

	return router, config
}
