package users

import (
	"log"
	"pkg/database"
)

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
