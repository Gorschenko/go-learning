package cars

import (
	"log/slog"
	"pkg/configs"
	"pkg/database"
)

type CarsRepository struct {
	Database *database.Db
	Config   *configs.Config
}

func NewCarsRepository(dependencies *database.RepositoryDependencies) *CarsRepository {
	needToMigrate := dependencies.Config.Database.Automigrate

	if needToMigrate {
		dependencies.Database.AutoMigrate(&database.Car{})
		slog.Debug("Car automigrate completed")
	}

	return &CarsRepository{
		Database: dependencies.Database,
	}
}
