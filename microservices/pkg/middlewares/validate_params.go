package middlewares

import (
	"context"
	"net/http"
	"pkg/api"
	"pkg/static"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateParams[DTO any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var params DTO

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
				err := api.
					NewInternalError(api.CodeBadRequest).
					WithMessage(err.Error())
				api.SendJSONError(w, err)
				return
			}
		}

		validate := validator.New()

		err := validate.Struct(params)

		if err != nil {
			err := api.
				NewInternalError(api.CodeBadRequest).
				WithMessage(err.Error())
			api.SendJSONError(w, err)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, static.ContextParamsKey, params)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
