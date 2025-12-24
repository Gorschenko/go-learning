package main

import (
	"log"
	"pkg/api"
	auth_api "pkg/api/auth"
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
	authApi := auth_api.NewAuthApi(&auth_api.AuthApiDependencies{
		Config:  config,
		HttpApi: httpApi,
	})

	ID, err := authApi.RegisterUser(&database.User{
		Email:    "123@123.com",
		Password: "123",
		Name:     "123",
	})

	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	log.Printf("User: %s", ID)
}
