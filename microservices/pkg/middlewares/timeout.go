package middlewares

import (
	"context"
	"net/http"
	"pkg/api"
	"pkg/errors"
	"pkg/logger"
	"time"
)

func TimeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Создаем контекст с таймаутом
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			logger := logger.GetLogger(ctx)
			defer cancel()

			// Создаем канал для отслеживания завершения обработки
			done := make(chan struct{})

			go func() {
				// Вызываем следующий обработчик
				next.ServeHTTP(w, r.WithContext(ctx))
				close(done)
			}()

			// Ждем либо завершения обработки, либо таймаута
			select {
			case <-done:
				// Обработка завершилась успешно
				return
			case <-ctx.Done():
				// Сработал таймаут
				if ctx.Err() == context.DeadlineExceeded {
					logger.Warn("TimeoutMiddleware", "Timeout", timeout)
					internalError := errors.NewInternalError(errors.CodeRequestTimeout)
					api.SendJSONError(w, internalError)
					return
				}
			}
		})
	}
}
