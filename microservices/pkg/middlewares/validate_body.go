package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"pkg/api"
	"pkg/static"

	"github.com/go-playground/validator/v10"
)

func ValidateBody[Body any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body Body

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			api.SendJSONError(w, errors.New(api.CodeBadRequest))
			return
		}

		validate := validator.New()
		err = validate.Struct(body)

		if err != nil {
			api.SendJSONError(w, errors.New(api.CodeBadRequest))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, static.ContextBodyKey, body)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
