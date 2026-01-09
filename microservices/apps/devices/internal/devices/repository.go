package devices

import (
	"log/slog"
	"pkg/database"
)

type DevicesRepository struct {
	database *database.Db
}

func NewDevicesRepository(dependencies *database.RepositoryDependencies) *DevicesRepository {
	needToMigrate := dependencies.Config.Database.Automigrate

	if needToMigrate {
		dependencies.Database.DB.AutoMigrate(&database.Device{})
		slog.Debug("Device automigrate completed")
	}

	return &DevicesRepository{
		database: dependencies.Database,
	}
}
