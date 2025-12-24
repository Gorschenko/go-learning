package users

import (
	"log"
	"pkg/configs"
	"pkg/database"
)

type UsersRepository struct {
	Database *database.Db
	Config   *configs.Config
}

type UsersRepositoryDependencies struct {
	Database *database.Db
	Config   *configs.Config
}

func NewUsersRepository(dependencies *UsersRepositoryDependencies) *UsersRepository {
	needToMigrate := dependencies.Config.Database.Automigrate
	if needToMigrate {
		dependencies.Database.DB.AutoMigrate(&database.User{})
		log.Println("User automigrate completed")
	}

	return &UsersRepository{
		Database: dependencies.Database,
	}
}

func (repository *UsersRepository) Create(user *database.User) (*database.User, error) {
	result := repository.Database.DB.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
