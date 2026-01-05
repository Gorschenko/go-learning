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

func (r *UsersRepository) UpdateOne(filters *users_api.UserFiltersDto, update *users_api.UserUpdateDto) (int64, error) {
	query := r.Database.DB.
		Model(&database.User{})
	query = r.prepareQuery(query, filters)

	result := query.Limit(1).UpdateColumns(update)

	return result.RowsAffected, result.Error
}

func (r *UsersRepository) GetOne(filters *users_api.UserFiltersDto) (*database.User, error) {
	var user database.User

	query := r.Database.DB.
		Model(&database.User{})
	query = r.prepareQuery(query, filters)

	result := query.First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UsersRepository) DeleteOne(filters *users_api.UserFiltersDto) (int64, error) {
	query := r.Database.DB.
		Model(&database.User{})
	query = r.prepareQuery(query, filters)

	result := query.Limit(1).Delete(&database.User{})

	return result.RowsAffected, result.Error
}

func (r *UsersRepository) prepareQuery(q *gorm.DB, filters *users_api.UserFiltersDto) (query *gorm.DB) {
	if filters.Email != "" {
		q = q.Where("email = ?", filters.Email)
	}

	if filters.ID != 0 {
		q = q.Where("id = ?", filters.ID)
	}

	return q
}
