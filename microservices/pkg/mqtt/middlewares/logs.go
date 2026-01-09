package mqtt_middlewares

import (
	"context"
	"pkg/logger"
	pkg_mqtt "pkg/mqtt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func LogsMiddleware(next pkg_mqtt.Handler) pkg_mqtt.Handler {
	return func(ctx context.Context, message mqtt.Message) {
		logger := logger.GetLogger(ctx)
		start := time.Now()

		next(ctx, message)

		logger.Info(
			"MqttLogsMiddleware",
			"Topic", message.Topic(),
			"Payload", message.Payload(),
			"Duration", time.Since(start),
		)
	}
}
