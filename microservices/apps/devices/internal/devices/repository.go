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

func (r *DevicesRepository) UpdateOne(serialNumber string, update DeviceUpdateDto) (int64, error) {
	result := r.database.DB.
		Model(&database.User{}).
		Where("serial_number = ?", serialNumber).
		Limit(1).
		UpdateColumns(update)

	return result.RowsAffected, result.Error
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

	result := r.database.DB.
		Model(&database.Device{}).
		Where("serial_number = ?", serialNumber).
		First(&device)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &device, nil
}
