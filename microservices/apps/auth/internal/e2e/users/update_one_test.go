package e2e_users

import (
	"auth/internal/e2e"
	"net/http"
	users_api "pkg/api/users"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestUpdateOneUser(t *testing.T) {
	ts, user := e2e.Setup(t)

	t.Run("Positive", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusOK), func(t *testing.T) {

			filters := users_api.UserFiltersDto{
				Email: user.Email,
			}

			update := users_api.UserUpdateDto{
				Name: gofakeit.Name(),
			}

			requestBody := users_api.UpdateOneRequestBodyDto{
				Filters: filters,
				Update:  update,
			}

			URL := ts.URL + users_api.UpdateOnePath
			var responseBody users_api.UpdateOneResponseBodyDto

			response, _ := resty.
				New().
				R().
				SetBody(requestBody).
				SetResult(&responseBody).
				Execute(users_api.UpdateOneMethod, URL)

			assert.Equal(t, http.StatusOK, response.StatusCode())
			assert.Equal(t, 1, responseBody.Count)
		})
	})

	t.Run("Negative", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusNotFound), func(t *testing.T) {
			filters := users_api.UserFiltersDto{
				Email: gofakeit.Name(),
			}

			update := users_api.UserUpdateDto{
				Name: gofakeit.Name(),
			}

			requestBody := users_api.UpdateOneRequestBodyDto{
				Filters: filters,
				Update:  update,
			}

			URL := ts.URL + users_api.UpdateOnePath
			var responseBody users_api.UpdateOneResponseBodyDto

			response, _ := resty.
				New().
				R().
				SetBody(requestBody).
				SetResult(&responseBody).
				Execute(users_api.UpdateOneMethod, URL)

			assert.Equal(t, http.StatusNotFound, response.StatusCode())
		})
	})
}
