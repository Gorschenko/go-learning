package main

import (
	"log"
	"pkg/api"
	"pkg/configs"
	"pkg/database"
)

func main() {
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

	user, err := authApi.RegisterUser(&database.User{
		Email:    "123@123.com",
		Password: "123",
		Name:     "123",
	})

	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	log.Printf("User: %s", user)
}
