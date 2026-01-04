package e2e

import (
	"net/http"
	users_api "pkg/api/users"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestDeleteOneUser(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusOK), func(t *testing.T) {
			user := CreateUser()

			requestBody := users_api.DeleteOneRequestBodyDto{
				ID: int(user.ID),
			}

			URL := testServer.URL + users_api.DeleteOnePath
			var responseBody users_api.DeleteOneResponseBodyDto

			response, _ := resty.
				New().
				R().
				SetHeader("Content-Type", "application/json").
				SetBody(requestBody).
				SetResult(&responseBody).
				Execute(users_api.DeleteOneMethod, URL)

			assert.Equal(t, http.StatusOK, response.StatusCode())
			assert.Equal(t, 1, responseBody.Count)
		})

	})
}
