package devices

import (
	"math/rand"
	"pkg/database"
)

func GetRandomDeviceStatus() database.DeviceStatus {
	num := rand.Intn(100)

	if num <= 10 {
		return database.DevicesStatusUnknown
	} else if num > 10 && num <= 50 {
		return database.DeviceStatusOnline
	} else {
		return database.DeviceStatusOffline
	}
}

func GetRandomDeviceType() database.DeviceType {
	num := rand.Intn(100)

	if num <= 10 {
		return database.DeviceTypeUnknown
	} else if num > 10 && num <= 50 {
		return database.DeviceTypeIgla
	} else {
		return database.DeviceTypeCompass
	}
}
