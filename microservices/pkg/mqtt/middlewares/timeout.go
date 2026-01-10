package mqtt_middlewares

import (
	"context"
	"pkg/logger"
	pkg_mqtt "pkg/mqtt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TimeoutMiddleware(timeout time.Duration) func(pkg_mqtt.Handler) pkg_mqtt.Handler {
	return func(next pkg_mqtt.Handler) pkg_mqtt.Handler {
		return func(ctx context.Context, message mqtt.Message) {

			// Создаем контекст с таймаутом
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			logger := logger.GetLogger(ctx)
			// Создаем канал для отслеживания завершения обработки
			done := make(chan struct{})

			go func() {
				// Вызываем следующий обработчик
				next(ctx, message)
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
					logger.Warn("MqttTimeoutMiddleware", "Timeout", timeout)
					return
				}
			}

		}
	}

}
