package app

import (
	"net/http"
	"pkg/configs"
	"pkg/database"
	"pkg/logger"
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

	router := http.NewServeMux()

	return router, config
}
