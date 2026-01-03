package users_e2e

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	auth_api "pkg/api/auth"
	users_api "pkg/api/users"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"
)

func TestGetOneUser(t *testing.T) {
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

			requestGetOneUserQuery := users_api.GetOneRequestQueryDto{
				Email: requestRegisterBody.Email,
			}
			requestGetOneQueryValues, _ := query.Values(requestGetOneUserQuery)
			URL = testServer.URL + users_api.GetOnePath
			URL = URL + "?" + requestGetOneQueryValues.Encode()

			responseGetOne, _ := http.Get(URL)
			responseGetOneBodyString, _ := io.ReadAll(responseGetOne.Body)
			var responseGetOneBody users_api.GetOneResponseBodyDto
			json.Unmarshal(responseGetOneBodyString, &responseGetOneBody)

			assert.Equal(t, http.StatusOK, responseGetOne.StatusCode)
			assert.Equal(t, requestRegisterBody.Email, responseGetOneBody.User.Email)
		})
	})

	t.Run("Negative", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusNotFound), func(t *testing.T) {
			requestQuery := users_api.GetOneRequestQueryDto{
				ID: gofakeit.Int(),
			}
			queryValues, _ := query.Values(requestQuery)
			URL := testServer.URL + users_api.GetOnePath
			URL = URL + "?" + queryValues.Encode()

			response, _ := http.Get(URL)
			responseBodyString, _ := io.ReadAll(response.Body)
			var responseBody users_api.GetOneResponseBodyDto
			json.Unmarshal(responseBodyString, &responseBody)

			assert.Equal(t, http.StatusNotFound, response.StatusCode)
		})
	})
}
