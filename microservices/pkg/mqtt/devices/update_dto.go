package mqtt_devices_api

type DeviceStatus string

const (
	TopicDevicesUpdated = "device.updated"
)

const (
	DeviceStatusOnline  DeviceStatus = "online"
	DeviceStatusOffline DeviceStatus = "offline"
)

type DeviceUpdateDto struct {
	Status DeviceStatus `json:"status"`
}
