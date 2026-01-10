package devices

import "pkg/database"

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
		createdDevice, _ := s.devicesRepository.CreateOne(&database.Device{
			SerialNumber: serialNumber,
			Status:       status,
		})
		device = createdDevice
	}

	return device
}
