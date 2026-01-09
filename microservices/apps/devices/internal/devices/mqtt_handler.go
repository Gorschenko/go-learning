package devices

import (
	"context"
	pkg_mqtt "pkg/mqtt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttDevicesHandler struct{}

func NewMqttDevicesHandler(mqttService *pkg_mqtt.MqttService) {
	handler := &MqttDevicesHandler{}

	mqttService.Subscribe("#", 0, handler.UpdateDevice())
}

func (h *MqttDevicesHandler) UpdateDevice() pkg_mqtt.HandlerFunc {
	return func(ctx context.Context, message mqtt.Message) {

	}
}
