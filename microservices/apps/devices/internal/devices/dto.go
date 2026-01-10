package devices

import "pkg/database"

type DeviceUpdateDto struct {
	Type   database.DeviceType   `json:"type" validate:"oneof=igla compass"`
	Status database.DeviceStatus `json:"status" validate:"required,oneof=online offline"`
}
