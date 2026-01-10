package devices

import (
	"context"
	pkg_mqtt "pkg/mqtt"
	mqtt_devices_api "pkg/mqtt/devices"
	mqtt_middlewares "pkg/mqtt/middlewares"
	"pkg/static"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMqttDevicesHandler(dependencies *MqttDevicesHandlerDependencies) {
	handler := &MqttDevicesHandler{
		mqttService:    dependencies.MqttService,
		devicesService: dependencies.DevicesService,
	}

	udpdateDeviceMiddlewares := mqtt_middlewares.CombineMiddlewares(
		mqtt_middlewares.CorrelationIdMiddleware,
		mqtt_middlewares.LogsMiddleware,
		mqtt_middlewares.TimeoutMiddleware(5*time.Second),
		mqtt_middlewares.ValidatePayload[mqtt_devices_api.DeviceUpdateDto],
	)
	handler.mqttService.Subscribe("#", 0,
		udpdateDeviceMiddlewares(handler.UpdateDeviceStatus()),
	)
}

func (h *MqttDevicesHandler) UpdateDeviceStatus() pkg_mqtt.Handler {
	return func(ctx context.Context, message mqtt.Message) {
		payload, _ := ctx.Value(static.ContextPayloadKey).(mqtt_devices_api.DeviceUpdateDto)
		topic := message.Topic()
		parts := strings.Split(topic, "/")
		serialNumber := parts[0]
		h.devicesService.UpdateDeviceStatus(serialNumber, payload.Status)
	}
}
