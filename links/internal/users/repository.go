package users

import "test/packages/db"

type UsersRepository struct {
	Database *db.Db
}

func NewUsersRepository(database *db.Db) *UsersRepository {
	return &UsersRepository{
		Database: database,
	}
}

func (repository *UsersRepository) Create(user *User) (*User, error) {
	result := repository.Database.DB.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repository *UsersRepository) FindByEmail(email string) (*User, error) {
	var user User

	result := repository.Database.DB.First(&user, "email = ?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
