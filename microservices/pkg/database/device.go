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
	SerialNumber string       `gorm:"index;not null"`
	Type         DeviceType   `gorm:"not null;type:ENUM('igla','compass','unknown');default:'unknown'"`
	Status       DeviceStatus `gorm:"not null;type:ENUM('online,'offline');default:'offline'"`
}
