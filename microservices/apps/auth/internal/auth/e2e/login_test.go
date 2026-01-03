package auth_e2e

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	auth_api "pkg/api/auth"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {

	t.Run("Positive", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusOK), func(t *testing.T) {
			URL := testServer.URL + auth_api.RegisterPath
			requestRegisterBody := auth_api.RegisterRequestBodyDto{
				Email:    gofakeit.Email(),
				Password: gofakeit.Password(false, false, true, false, false, 5),
				Name:     gofakeit.Name(),
			}
			requestRegisterBodyString, _ := json.Marshal(&requestRegisterBody)
			responseRegister, _ := http.Post(URL, "application/json", bytes.NewReader(requestRegisterBodyString))

			assert.NotNil(t, responseRegister)

			requestLoginBody, _ := json.Marshal(&auth_api.LoginRequestBodyDto{
				Email:    requestRegisterBody.Email,
				Password: requestRegisterBody.Password,
			})

			URL = testServer.URL + auth_api.LoginPath
			responseLogin, _ := http.Post(URL, "application/json", bytes.NewReader(requestLoginBody))
			responseBodyLoginString, _ := io.ReadAll(responseLogin.Body)
			var responseBody auth_api.RegisterResponseBodyDto
			json.Unmarshal(responseBodyLoginString, &responseBody)

			assert.Equal(t, http.StatusOK, responseLogin.StatusCode)
			assert.NotEmpty(t, responseBody.Token)
			assert.False(t, responseBody.ExpirationTime.IsZero())
		})
	})

	t.Run("Negative", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusNotFound), func(t *testing.T) {
			URL := testServer.URL + auth_api.LoginPath
			requestBody, _ := json.Marshal(&auth_api.LoginRequestBodyDto{
				Email:    gofakeit.Email(),
				Password: gofakeit.Password(false, false, true, false, false, 5),
			})
			response, _ := http.Post(URL, "application/json", bytes.NewReader(requestBody))

			assert.Equal(t, http.StatusNotFound, response.StatusCode)
		})
	})
}
