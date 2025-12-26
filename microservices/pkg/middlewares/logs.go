package middlewares

import (
	"log/slog"
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

		slog.Info(
			"[LogsMiddleware]",
			"Method", r.Method,
			"Path", r.URL.Path,
			"Status", wrapper.StatusCode,
			"Duration", time.Since(start),
		)
	})
}
