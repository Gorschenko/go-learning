package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LogsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r)

		log.Printf(
			"%s - Method: %s, Path: %s, Status: %d, Duration: %v",
			"[LogMiddleware]",
			r.Method,
			r.URL.Path,
			wrapper.StatusCode,
			time.Since(start),
		)
	})
}
