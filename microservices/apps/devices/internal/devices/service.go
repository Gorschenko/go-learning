package devices

import (
	"pkg/database"
)

func NewDevicesService(dependencies *DevicesServiceDependencies) *DevicesService {
	return &DevicesService{
		devicesRepository: dependencies.DevicesRepository,
	}
}

func (s *DevicesService) UpdateDeviceStatus(serialNumber string, status database.DeviceStatus) *database.Device {
	var device *database.Device

	existedDevice, _ := s.devicesRepository.GetOne(serialNumber)

	if existedDevice != nil {
		device = existedDevice
	}

	if existedDevice == nil {
		dataToCreate := database.Device{
			Type:         GetRandomDeviceType(),
			SerialNumber: serialNumber,
			Status:       status,
		}

		createdDevice, _ := s.devicesRepository.CreateOne(&dataToCreate)
		device = createdDevice
	}

	return device
}
