package devices

import (
	"context"
	"fmt"
	pkg_mqtt "pkg/mqtt"
	mqtt_devices_api "pkg/mqtt/devices"
	mqtt_middlewares "pkg/mqtt/middlewares"
	"pkg/static"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttDevicesHandler struct{}

func NewMqttDevicesHandler(mqttService *pkg_mqtt.MqttService) {
	handler := &MqttDevicesHandler{}

	udpdateDeviceMiddlewares := mqtt_middlewares.CombineMiddlewares(
		mqtt_middlewares.CorrelationIdMiddleware,
		mqtt_middlewares.LogsMiddleware,
		mqtt_middlewares.TimeoutMiddleware(5*time.Second),
		mqtt_middlewares.ValidatePayload[mqtt_devices_api.DeviceUpdateDto],
	)
	mqttService.Subscribe("#", 0,
		udpdateDeviceMiddlewares(handler.UpdateDevice()),
	)
}

func (h *MqttDevicesHandler) UpdateDevice() pkg_mqtt.Handler {
	return func(ctx context.Context, message mqtt.Message) {
		payload, _ := ctx.Value(static.ContextPayloadKey).(mqtt_devices_api.DeviceUpdateDto)
		fmt.Printf("Payload from handler %+v\n", payload)
	}
}
