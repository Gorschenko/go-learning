package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"pkg/api"
	"pkg/static"

	"github.com/go-playground/validator/v10"
)

func ValidateBody[T any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body T

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			api.SendJSON(w, static.ErorInvalidJSON, http.StatusBadRequest)
			return
		}

		validate := validator.New()
		err = validate.Struct(body)

		if err != nil {
			api.SendJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), static.ContextBodyKey, body)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
