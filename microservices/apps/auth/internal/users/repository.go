package users

import (
	"errors"
	"log/slog"
	users_api "pkg/api/users"
	"pkg/database"

	"gorm.io/gorm"
)

type UsersRepository struct {
	Database *database.Db
}

func NewUsersRepository(dependencies *database.RepositoryDependencies) *UsersRepository {
	needToMigrate := dependencies.Config.Database.Automigrate

	if needToMigrate {
		dependencies.Database.DB.AutoMigrate(&database.User{})
		slog.Debug("User automigrate completed")
	}

	return &UsersRepository{
		Database: dependencies.Database,
	}
}

func (r *UsersRepository) Create(user *database.User) (*database.User, error) {
	result := r.Database.DB.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// func (r *UsersRepository) Update(user *database.User) (*database.User, error) {

// }

func (r *UsersRepository) GetOne(filters *users_api.UserFiltersDto) (*database.User, error) {
	var user database.User

	query := r.Database.DB.
		Model(&database.User{})

	if filters.Email != "" {
		query = query.Where("email = ?", filters.Email)
	}

	if filters.ID != 0 {
		query = query.Where("id = ?", filters.ID)
	}

	result := query.First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UsersRepository) DeleteOne(filters *users_api.UserFiltersDto) (int, error) {
	query := r.Database.DB.
		Model(&database.User{})

	if filters.Email != "" {
		query = query.Where("email = ?", filters.Email)
	}

	if filters.ID != 0 {
		query = query.Where("id = ?", filters.ID)
	}

	result := query.Limit(1).Delete(&database.User{})

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}
