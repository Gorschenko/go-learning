package e2e

import (
	"auth/internal/app"
	"net/http/httptest"
	"pkg/api"
	auth_api "pkg/api/auth"
	users_api "pkg/api/users"
	"pkg/database"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-resty/resty/v2"
)

func Setup(t *testing.T) (*httptest.Server, *database.User) {
	app, _ := app.GetApp("../../../../../config.json")
	ts := httptest.NewServer(app)
	user := CreateUser(ts.URL)

	t.Cleanup(func() {
		defer ts.Close()

		filters := users_api.UserFiltersDto{
			ID: int(user.ID),
		}
		defer DeleteUser(ts.URL, &filters)

	})

	return ts, user
}

func CreateUser(URL string) *database.User {
	requestBody := auth_api.RegisterRequestBodyDto{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(false, false, true, false, false, 5),
		Name:     gofakeit.Name(),
	}

	httpApi := api.HttpApi{
		Client: resty.New(),
	}
	authHttpApi := auth_api.
		NewAuthApi(&auth_api.AuthApiDependencies{
			HttpApi: &httpApi,
		}).
		SetBaseURL(URL)

	response, _ := authHttpApi.RegisterUser(&requestBody)

	return response.User
}

func DeleteUser(URL string, filters *users_api.UserFiltersDto) int {
	httpApi := api.HttpApi{
		Client: resty.New(),
	}
	usersHttpApi := users_api.
		NewUsersApi(&users_api.UsersApiDependencies{
			HttpApi: &httpApi,
		}).
		SetBaseURL(URL)

	response, _ := usersHttpApi.DeleteOneUser(filters)

	return response.Count
}
