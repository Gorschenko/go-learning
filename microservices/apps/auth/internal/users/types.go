package users

import (
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
