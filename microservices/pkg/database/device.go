package database

import "gorm.io/gorm"

type DeviceType string

const (
	DeviceTypeIgla    = "igla"
	DeviceTypeCompass = "compass"
	DeviceTypeUnknown = "unknown"
)

type DeviceStatus string

const (
	DeviceStatusOnline  DeviceStatus = "online"
	DeviceStatusOffline DeviceStatus = "offline"
	// Для обработки ошибок
	DevicesStatusUnknown DeviceStatus = "unknown"
)

type Device struct {
	gorm.Model
	SerialNumber string       `gorm:"uniqueIndex;not null"`
	Type         DeviceType   `gorm:"not null;type:varchar(20);default:'unknown'"`
	Status       DeviceStatus `gorm:"not null;type:varchar(20);default:'offline'"`
}
