package devices

import (
	"errors"
	"log/slog"
	"pkg/database"

	"gorm.io/gorm"
)

func NewDevicesRepository(dependencies *database.RepositoryDependencies) *DevicesRepository {
	needToMigrate := dependencies.Config.Database.Automigrate

	if needToMigrate {
		err := dependencies.Database.DB.AutoMigrate(&database.Device{})
		if err != nil {
			panic(err)
		}
		slog.Debug("Device automigrate completed")
	}

	return &DevicesRepository{
		database: dependencies.Database,
	}
}

func (r *DevicesRepository) CreateOne(device *database.Device) (*database.Device, error) {
	result := r.database.DB.
		Model(&database.Device{}).
		Create(device)

	if result.Error != nil {
		return nil, result.Error
	}

	return device, nil
}

func (r *DevicesRepository) GetOne(serialNumber string) (*database.Device, error) {
	var device database.Device

	query := r.database.DB.
		Model(&database.Device{}).
		Where("serial_number = ?", serialNumber)

	result := query.First(&device)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &device, nil
}
