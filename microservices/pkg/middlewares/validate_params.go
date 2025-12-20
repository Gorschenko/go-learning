package middlewares

import (
	"context"
	"net/http"
	"pkg/static"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateParams[T any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var params T

		v := reflect.ValueOf(&params).Elem()
		t := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			paramName := field.Tag.Get("param")

			if paramName == "" {
				paramName = strings.ToLower(field.Name)
			}

			paramValue := r.PathValue(paramName)

			if paramValue == "" {
				continue
			}

			err := setFieldFromString(v.Field(i), paramValue)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		validate := validator.New()

		err := validate.Struct(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), static.ContextParamsKey, params)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
