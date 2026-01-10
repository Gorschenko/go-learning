package e2e_auth

import (
	"auth/internal/e2e"
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
	ts, user := e2e.Setup(t)

	t.Run("Positive", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusOK), func(t *testing.T) {

			requestBody, _ := json.Marshal(&auth_api.LoginRequestBodyDto{
				Email:    user.Email,
				Password: user.Password,
			})

			URL := ts.URL + auth_api.LoginPath
			response, _ := http.Post(URL, "application/json", bytes.NewReader(requestBody))
			responseBodyString, _ := io.ReadAll(response.Body)
			var responseBody auth_api.LoginResponseBodyDto
			json.Unmarshal(responseBodyString, &responseBody)

			assert.Equal(t, http.StatusOK, response.StatusCode)
			assert.NotEmpty(t, responseBody.Token)
			assert.False(t, responseBody.ExpirationTime.IsZero())

		})
	})

	t.Run("Negative", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusNotFound), func(t *testing.T) {
			URL := ts.URL + auth_api.LoginPath
			requestBody, _ := json.Marshal(&auth_api.LoginRequestBodyDto{
				Email:    gofakeit.Email(),
				Password: gofakeit.Password(false, false, true, false, false, 5),
			})
			response, _ := http.Post(URL, "application/json", bytes.NewReader(requestBody))

			assert.Equal(t, http.StatusNotFound, response.StatusCode)
		})
	})
}
