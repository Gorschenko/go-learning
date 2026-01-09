package devices

import (
	"context"
	"fmt"
	pkg_mqtt "pkg/mqtt"
	mqtt_devices_api "pkg/mqtt/devices"
	"pkg/static"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttDevicesHandler struct{}

func NewMqttDevicesHandler(mqttService *pkg_mqtt.MqttService) {
	handler := &MqttDevicesHandler{}

	mqttService.Subscribe("#", 0, pkg_mqtt.ValidatePayload[mqtt_devices_api.DeviceUpdateDto](handler.UpdateDevice()))
}

func (h *MqttDevicesHandler) UpdateDevice() pkg_mqtt.HandlerFunc {
	return func(ctx context.Context, message mqtt.Message) {
		payload, _ := ctx.Value(static.ContextPayloadKey).(mqtt_devices_api.DeviceUpdateDto)
		fmt.Printf("Payload %+v\n", payload)
	}
}
