package mqtt_middlewares

import (
	"context"
	"encoding/json"
	"pkg/logger"
	pkg_mqtt "pkg/mqtt"
	"pkg/static"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-playground/validator/v10"
)

func ValidatePayload[Payload any](next pkg_mqtt.Handler) pkg_mqtt.Handler {
	return func(ctx context.Context, message mqtt.Message) {
		var payload Payload

		err := json.Unmarshal(message.Payload(), &payload)

		if err != nil {
			logger := logger.GetLogger(ctx)
			logger.Debug(
				"ValidatePayloadMiddleware",
				"Failed to unmarshal payload", err.Error(),
				"Topic", message.Topic(),
				"Payload", string(message.Payload()),
			)
			return
		}

		validate := validator.New()
		err = validate.Struct(payload)

		if err != nil {
			logger := logger.GetLogger(ctx)
			logger.Debug(
				"ValidatePayloadMiddleware",
				"Validation failed", err.Error(),
				"Topic", message.Topic(),
				"Payload", string(message.Payload()),
			)
			return
		}

		ctx = context.WithValue(ctx, static.ContextPayloadKey, payload)

		next(ctx, message)
	}
}
