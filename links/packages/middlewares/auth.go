package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"test/configs"
	"test/packages/jwt"
)

type contextKey struct{}

var (
	ContextEmailKey = contextKey{}
)

func IsAuthenticated(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		fmt.Println(isValid)

		r.Context()
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		new_req := r.WithContext(ctx)

		next.ServeHTTP(w, new_req)
	})
}
