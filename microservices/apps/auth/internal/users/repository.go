package users

import (
	"log/slog"
	"pkg/database"
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

func (r *UsersRepository) FindByEmail(email string) (*database.User, error) {
	var user database.User

	result := r.Database.DB.First(&user, "email = ?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UsersRepository) FindOne(filters FindOneUserFilters) (*database.User, error) {
	var user database.User

	query := r.Database.DB.
		Model(&database.User{})

	if filters.Email != "" {
		query = query.Where("email = ?", filters.Email)
	}

	if filters.UserID != 0 {
		query = query.Where("user_id = ?", filters.UserID)
	}

	result := query.First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
