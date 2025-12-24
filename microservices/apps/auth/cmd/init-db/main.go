package main

import (
	"log"
	"pkg/api"
	auth_api "pkg/api/auth"
	"pkg/configs"
	"pkg/database"
	"time"

	"github.com/brianvoe/gofakeit/v7"
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

	usersCount := config.InitDatabase.Users.Count
	status := &OperationStatus{
		total:    usersCount,
		fulfiled: 0,
		rejected: 0,
	}

	for i := 0; i < usersCount; i++ {
		user := &database.User{
			Email:    gofakeit.Email(),
			Password: gofakeit.Password(false, false, true, false, false, 5),
			Name:     gofakeit.Name(),
		}

		ID, err := authApi.RegisterUser(user)

		if err != nil {
			status.rejected++
			log.Printf("User created with ID %d: ", ID)
		} else {
			status.fulfiled++
		}

		time.Sleep(1 * time.Second)
	}

	log.Printf("Body: %+v\n", status)
}
