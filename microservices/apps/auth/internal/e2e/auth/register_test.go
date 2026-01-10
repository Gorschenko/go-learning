package e2e_auth

import (
	"auth/internal/e2e"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	auth_api "pkg/api/auth"
	users_api "pkg/api/users"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	ts, _ := e2e.Setup(t)

	t.Run("Positive", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusOK), func(t *testing.T) {
			user := e2e.CreateUser(ts.URL)

			fmt.Printf("Created user %+v\n", user)
			assert.NotEmpty(t, user)

			filters := users_api.UserFiltersDto{
				ID: int(user.ID),
			}

			count := e2e.DeleteUser(ts.URL, &filters)

			fmt.Printf("Deleted users count %d", count)
		})
	})

	t.Run("Negative", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusBadRequest), func(t *testing.T) {
			URL := ts.URL + auth_api.RegisterPath
			requestBodyString, _ := json.Marshal(&auth_api.RegisterRequestBodyDto{
				Password: gofakeit.Password(false, false, true, false, false, 5),
				Name:     gofakeit.Name(),
			})
			response, _ := http.Post(URL, "application/json", bytes.NewReader(requestBodyString))

			assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		})
	})
}
