package devices

import (
	"pkg/database"
	"pkg/mqtt"
)

type MqttDevicesHandlerDependencies struct {
	MqttService    *mqtt.MqttService
	DevicesService *DevicesService
}

type MqttDevicesHandler struct {
	mqttService    *mqtt.MqttService
	devicesService *DevicesService
}

type DevicesServiceDependencies struct {
	DevicesRepository *DevicesRepository
}

type DevicesService struct {
	devicesRepository *DevicesRepository
}

type DevicesRepository struct {
	database *database.Db
}
