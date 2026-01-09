package mqtt_middlewares

import (
	"context"
	"pkg/logger"
	pkg_mqtt "pkg/mqtt"
	"pkg/static"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

func CorrelationIdMiddleware(next pkg_mqtt.Handler) pkg_mqtt.Handler {
	return func(ctx context.Context, message mqtt.Message) {
		logger := logger.GetLogger(ctx)

		/*
			Нужно исправить:
			Нужно добавлять correlationId, если он передан в событие
		*/

		correlationId := uuid.New().String()

		logger.Debug(
			"MqttCorrelationIDMiddleware",
			"CorelationID", correlationId,
		)

		ctx = context.WithValue(ctx, static.ContextCorrelationID, correlationId)

		next(ctx, message)
	}
}
