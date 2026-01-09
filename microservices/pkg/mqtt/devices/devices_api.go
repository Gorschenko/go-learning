package mqtt_devices_api

import (
	"context"
	"pkg/mqtt"
)

type DevicesDependencies struct {
	MqttService *mqtt.MqttService
}

type DevicesApi struct {
	mqttSevice *mqtt.MqttService
}

func NewDevicesApi(dependencies *DevicesDependencies) *DevicesApi {
	return &DevicesApi{
		mqttSevice: dependencies.MqttService,
	}
}

func (api *DevicesApi) SendUpdateDeviceEvent(ctx context.Context, serialNumber string, update *DeviceUpdateDto) error {
	topic := serialNumber + "/" + TopicDevicesUpdated
	err := api.mqttSevice.Publish(ctx, topic, 0, update)

	if err != nil {
		return err
	}

	return nil
}
