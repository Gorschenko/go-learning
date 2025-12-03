package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"test/configs"
	"test/packages/jwt"
	"test/packages/response"
)

type contextKey struct{}

var (
	ContextEmailKey = contextKey{}
)

func IsAuthenticated(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			response.Json(w, nil, http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)

		if !isValid {
			response.Json(w, nil, http.StatusUnauthorized)
			return
		}
		fmt.Println(isValid)

		r.Context()
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		new_req := r.WithContext(ctx)

		next.ServeHTTP(w, new_req)
	})
}
