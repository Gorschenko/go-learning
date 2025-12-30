package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"pkg/api"
	"pkg/errors"
	"pkg/static"

	"github.com/go-playground/validator/v10"
)

func ValidateBody[DTO any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body DTO

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			err := errors.
				NewInternalError(errors.CodeBadRequest).
				WithMessage(err.Error())
			api.SendJSONError(w, err)
			return
		}

		validate := validator.New()
		err = validate.Struct(body)

		if err != nil {
			err := errors.
				NewInternalError(errors.CodeBadRequest).
				WithMessage(err.Error())
			api.SendJSONError(w, err)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, static.ContextBodyKey, body)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
