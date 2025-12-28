package middlewares

import (
	"net/http"
	"pkg/logger"
	"time"
)

func LogsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logger.GetLogger(r.Context())

		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r.WithContext(ctx))

		logger.Info(
			"LogsMiddleware",
			"Method", r.Method,
			"Path", r.URL.Path,
			"Status", wrapper.StatusCode,
			"Duration", time.Since(start),
		)
	})
}
