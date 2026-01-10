package e2e

import (
	"auth/internal/app"
	"net/http"
	"net/http/httptest"
	"os"
	"pkg/api"
	auth_api "pkg/api/auth"
	"pkg/database"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-resty/resty/v2"
)

var (
	testApp    http.Handler
	testServer *httptest.Server
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	app, _ := app.GetApp("../../../../config.json")
	testApp = app
	ts := httptest.NewServer(testApp)
	testServer = ts
}

func teardown() {
	defer testServer.Close()
}

func RegisterUser() *database.User {
	requestBody := auth_api.RegisterRequestBodyDto{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(false, false, true, false, false, 5),
		Name:     gofakeit.Name(),
	}

	httpApi := api.HttpApi{
		Client: resty.New(),
	}
	authHttpApi := auth_api.NewAuthApi(&auth_api.AuthApiDependencies{
		HttpApi: &httpApi,
	}).SetBaseURL(testServer.URL)

	response, _ := authHttpApi.RegisterUser(&requestBody)

	return response.User
}
