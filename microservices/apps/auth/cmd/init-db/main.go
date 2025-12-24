package main

import (
	"log"
	"pkg/api"
	auth_api "pkg/api/auth"
	"pkg/configs"
	"pkg/database"
	"sync"
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

	result := make(chan bool, usersCount)
	var wg sync.WaitGroup

	for i := 0; i < usersCount; i++ {
		wg.Add(1)

		go func(userNumber int) {
			defer wg.Done()

			user := &database.User{
				Email:    gofakeit.Email(),
				Password: gofakeit.Password(false, false, true, false, false, 5),
				Name:     gofakeit.Name(),
			}

			ID, err := authApi.RegisterUser(user)

			if err != nil {
				log.Printf("Error creating user %d: %v", userNumber, err)

				result <- false
			} else {
				log.Printf("User %d created with ID %d: ", userNumber, ID)

				result <- true
			}

			time.Sleep(1 * time.Second)
		}(i)
	}

	// Ждем завершения всех горутин и закрываем канал
	go func() {
		wg.Wait()
		close(result)
	}()

	// Обрабатываем результаты
	for success := range result {
		if success {
			status.fulfiled++
		} else {
			status.rejected++
		}
	}

	log.Printf("Result: %+v\n", status)
}
