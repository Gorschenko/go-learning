package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	auth_api "pkg/api/auth"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusOK), func(t *testing.T) {
			user := RegisterUser()
			assert.NotEmpty(t, user)
		})
	})

	t.Run("Negative", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusBadRequest), func(t *testing.T) {
			URL := testServer.URL + auth_api.RegisterPath
			requestBodyString, _ := json.Marshal(&auth_api.RegisterRequestBodyDto{
				Password: gofakeit.Password(false, false, true, false, false, 5),
				Name:     gofakeit.Name(),
			})
			response, _ := http.Post(URL, "application/json", bytes.NewReader(requestBodyString))

			assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		})
	})
}
