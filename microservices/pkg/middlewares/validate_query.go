package middlewares

import (
	"context"
	"net/http"
	"pkg/api"
	"pkg/errors"
	"pkg/static"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateQuery[DTO any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var query DTO

		v := reflect.ValueOf(&query).Elem()
		t := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			queryName := field.Tag.Get("query")

			if queryName == "" {
				queryName = strings.ToLower(field.Name)
			}

			queryValue := r.URL.Query().Get(queryName)

			if queryValue == "" {
				continue
			}

			err := setFieldFromString(v.Field(i), queryValue)

			if err != nil {
				err := errors.
					NewInternalError(errors.CodeBadRequest).
					WithMessage(err.Error())
				api.SendJSONError(w, err)
				return
			}
		}

		validate := validator.New()

		err := validate.Struct(query)

		if err != nil {
			err := errors.
				NewInternalError(errors.CodeBadRequest).
				WithMessage(err.Error())
			api.SendJSONError(w, err)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, static.ContextQueryKey, query)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
