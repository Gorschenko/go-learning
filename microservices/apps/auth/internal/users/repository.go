package users

import "pkg/database"

type UsersRepository struct {
	Database *database.Db
}

type UsersRepositoryDependencies struct {
	Database *database.Db
}

func NewUsersRepository(dependencies *UsersRepositoryDependencies) *UsersRepository {
	return &UsersRepository{
		Database: dependencies.Database,
	}
}
