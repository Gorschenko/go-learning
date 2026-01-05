package e2e

import (
	"encoding/json"
	"io"
	"net/http"
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
			user := CreateUser()

			requestQuery := users_api.UserFiltersDto{
				Email: user.Email,
			}
			requestQueryValues, _ := query.Values(requestQuery)
			URL := testServer.URL + users_api.GetOnePath
			URL = URL + "?" + requestQueryValues.Encode()

			response, _ := http.Get(URL)
			responseBodyString, _ := io.ReadAll(response.Body)
			var responseBody users_api.GetOneResponseBodyDto
			json.Unmarshal(responseBodyString, &responseBody)

			assert.Equal(t, http.StatusOK, response.StatusCode)
			assert.Equal(t, user.Email, responseBody.User.Email)
		})
	})

	t.Run("Negative", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusNotFound), func(t *testing.T) {
			requestQuery := users_api.UserFiltersDto{
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
