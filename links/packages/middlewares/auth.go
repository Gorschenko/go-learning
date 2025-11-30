package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		fmt.Println(token)
		next.ServeHTTP(w, r)
	})
}
