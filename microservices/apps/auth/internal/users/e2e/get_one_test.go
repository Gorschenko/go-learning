package users_e2e

import (
	"net/http"
	"strconv"
	"testing"
)

func TestGetOneUser(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {

	})

	t.Run("Negative", func(t *testing.T) {
		t.Run(strconv.Itoa(http.StatusNotFound), func(t *testing.T) {
			// requestBodyDto := users_api.GetOneRequestQueryDto{
			// 	UserID: gofakeit.Int(),
			// }
			// requestBody, _ := json.Marshal(requestBodyDto)

			// URL := testServer.URL + users_api.GetOnePath

			// responseBody, _ := http.Get(URL, "application/json", bytes.NewReader(requestBody))
		})
	})
}
