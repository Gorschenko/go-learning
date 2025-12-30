package middlewares

import (
	"bytes"
	"io"
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

		var stringRequestBody string

		if r.Body != nil && r.Method != "GET" && r.Method != "HEAD" {
			bodyBytes, err := io.ReadAll(r.Body)
			if err == nil {
				stringRequestBody = string(bodyBytes)
				// Восстанавливаем тело для последующих обработчиков
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		next.ServeHTTP(wrapper, r.WithContext(ctx))

		logger.Info(
			"LogsMiddleware",
			"Method", r.Method,
			"Path", r.URL.Path,
			"Status", wrapper.StatusCode,
			"Duration", time.Since(start),
			"RequestBody", stringRequestBody,
			"RequestQuery", r.URL.Query(),
			"ResponseBody", wrapper.ResponseBody,
		)
	})
}
