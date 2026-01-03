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
			requestRegisterDTO := auth_api.RegisterRequestBodyDto{
				Email:    gofakeit.Email(),
				Password: gofakeit.Password(false, false, true, false, false, 5),
				Name:     gofakeit.Name(),
			}
			requestRegisterBody, _ := json.Marshal(&requestRegisterDTO)

			URL := testServer.URL + auth_api.RegisterPath
			responseRegister, _ := http.Post(URL, "application/json", bytes.NewReader(requestRegisterBody))

			assert.NotNil(t, responseRegister)

			requestLoginBody, _ := json.Marshal(&auth_api.LoginRequestBodyDto{
				Email:    requestRegisterDTO.Email,
				Password: requestRegisterDTO.Password,
			})

			URL = testServer.URL + auth_api.LoginPath

			responseLogin, _ := http.Post(URL, "application/json", bytes.NewReader(requestLoginBody))

			assert.Equal(t, http.StatusOK, responseLogin.StatusCode)

			responseBodyString, _ := io.ReadAll(responseLogin.Body)

			var responseBody auth_api.RegisterResponseBodyDto
			json.Unmarshal(responseBodyString, &responseBody)

			assert.NotEmpty(t, responseBody.Token)
			assert.False(t, responseBody.ExpirationTime.IsZero())
		})
	})

	t.Run("Negative", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusNotFound), func(t *testing.T) {
			requestBody, _ := json.Marshal(&auth_api.LoginRequestBodyDto{
				Email:    gofakeit.Email(),
				Password: gofakeit.Password(false, false, true, false, false, 5),
			})

			URL := testServer.URL + auth_api.LoginPath
			response, _ := http.Post(URL, "application/json", bytes.NewReader(requestBody))

			assert.Equal(t, http.StatusNotFound, response.StatusCode)
		})
	})
}
