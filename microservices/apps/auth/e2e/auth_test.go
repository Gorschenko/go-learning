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
)

func TestRegisterUserSuccess(t *testing.T) {
	app, _ := app.GetApp("../../../config.json")
	ts := httptest.NewServer(app)
	defer ts.Close()

	requestBody, _ := json.Marshal(&auth_api.RegisterRequestBodyDto{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(false, false, true, false, false, 5),
		Name:     gofakeit.Name(),
	})

	URL := ts.URL + auth_api.AuthRegisterPath
	response, err := http.Post(URL, "application/json", bytes.NewReader(requestBody))

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected %d got %d", http.StatusOK, response.StatusCode)
	}

	responseBodyString, _ := io.ReadAll(response.Body)
	var responseBody auth_api.RegisterResponseBodyDto
	json.Unmarshal(responseBodyString, &responseBody)

	if responseBody.Token == "" {
		t.Error("Expected token in response")
	}

	if responseBody.ExpirationTime.IsZero() {
		t.Error("Expected expiration time in response")
	}
}
