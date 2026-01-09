package mqtt_devices_api

type DeviceStatus string

const (
	TopicDevicesUpdated = "device.updated"
)

const (
	DeviceStatusOnline  DeviceStatus = "online"
	DeviceStatusOffline DeviceStatus = "offline"
	// Для обработки ошибок
	DevicesStatusError DeviceStatus = "error"
)

type DeviceUpdateDto struct {
	Status  DeviceStatus `json:"status" validate:"required"`
	Status2 string       `json:"status2" validate:"required"`
}
