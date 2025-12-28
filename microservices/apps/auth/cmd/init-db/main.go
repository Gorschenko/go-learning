package main

import (
	"log/slog"
	"pkg/api"
	auth_api "pkg/api/auth"
	"pkg/configs"
	"pkg/database"
	"pkg/logger"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {
	config, err := configs.LoadConfig("../../config.json")
	if err != nil {
		panic(err)
	}

	logger.SetupLogger(&logger.LoggerServiceDependencies{
		Config: config,
	})

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
				slog.Error(
					"Error creating user",
					"UserNumber",
					userNumber,
					"Error",
					err,
				)

				result <- false
			} else {
				slog.Info(
					"User created",
					"UserNumber",
					userNumber,
					"ID",
					ID,
				)

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
	slog.Info(
		"Result",
		"Status",
		status,
	)
}
