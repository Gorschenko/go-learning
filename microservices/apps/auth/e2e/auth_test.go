package e2e

import (
	"auth/internal/app"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	auth_api "pkg/api/auth"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUserPositive(t *testing.T) {
	app, _ := app.GetApp("../../../config.json")
	ts := httptest.NewServer(app)
	defer ts.Close()

	requestBody, _ := json.Marshal(&auth_api.RegisterRequestBodyDto{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(false, false, true, false, false, 5),
		Name:     gofakeit.Name(),
	})

	URL := ts.URL + auth_api.AuthRegisterPath
	response, _ := http.Post(URL, "application/json", bytes.NewReader(requestBody))

	assert.Equal(t, http.StatusOK, response.StatusCode)

	responseBodyString, _ := io.ReadAll(response.Body)

	var responseBody auth_api.RegisterResponseBodyDto
	json.Unmarshal(responseBodyString, &responseBody)

	assert.NotEmpty(t, responseBody.Token)
	assert.False(t, responseBody.ExpirationTime.IsZero())
}

func TestRegisterUserNegative(t *testing.T) {
	app, _ := app.GetApp("../../../config.json")
	ts := httptest.NewServer(app)
	defer ts.Close()

	requestBody, _ := json.Marshal(&auth_api.RegisterRequestBodyDto{
		Password: gofakeit.Password(false, false, true, false, false, 5),
		Name:     gofakeit.Name(),
	})

	URL := ts.URL + auth_api.AuthRegisterPath
	response, _ := http.Post(URL, "application/json", bytes.NewReader(requestBody))

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}
