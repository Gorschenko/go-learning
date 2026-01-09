package mqtt_devices_api

import "pkg/database"

const (
	TopicDevicesUpdated = "device.updated"
)

type DeviceUpdateDto struct {
	Status database.DeviceStatus `json:"status" validate:"required,oneof=online offline"`
}
