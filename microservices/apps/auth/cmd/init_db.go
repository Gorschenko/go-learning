package main

import (
	"pkg/api"
	"pkg/configs"
)

func init_db() {
	config, err := configs.LoadConfig("../../config.json")
	if err != nil {
		panic(err)
	}

	// api
	httpApi := api.NewHttpApi(&api.HttpApiDependencies{
		Config: config,
	})
	authApi := api.NewAuthApi(&api.AuthApiDependencies{
		Config:  config,
		HttpApi: httpApi,
	})
}
