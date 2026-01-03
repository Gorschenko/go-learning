package e2e

import (
	"auth/internal/app"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	auth_api "pkg/api/auth"
	"pkg/database"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
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

func CreateUser() *database.User {
	URL := testServer.URL + auth_api.RegisterPath
	requestBodyDto := auth_api.RegisterRequestBodyDto{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(false, false, true, false, false, 5),
		Name:     gofakeit.Name(),
	}
	requestBodyString, _ := json.Marshal(&requestBodyDto)
	response, _ := http.Post(URL, "application/json", bytes.NewReader(requestBodyString))
	responseBodyString, _ := io.ReadAll(response.Body)
	var responseBody auth_api.RegisterResponseBodyDto

	json.Unmarshal(responseBodyString, &responseBody)

	return responseBody.User
}
