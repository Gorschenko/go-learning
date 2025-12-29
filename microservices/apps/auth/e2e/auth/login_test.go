package auth_e2e

import (
	"auth/internal/app"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	auth_api "pkg/api/auth"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestLoginUserNegative(t *testing.T) {
	app, _ := app.GetApp("../../../../config.json")
	ts := httptest.NewServer(app)
	defer ts.Close()

	requestBody, _ := json.Marshal(&auth_api.LoginRequestBodyDto{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(false, false, true, false, false, 5),
	})

	URL := ts.URL + auth_api.AuthLoginPath
	response, _ := http.Post(URL, "application/json", bytes.NewReader(requestBody))

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}
