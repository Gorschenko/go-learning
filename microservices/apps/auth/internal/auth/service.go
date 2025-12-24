package auth

import (
	"auth/internal/users"
	"pkg/database"
)

type AuthServiceDependencies struct {
	UsersRepository *users.UsersRepository
}

type AuthService struct {
	UsersRespository *users.UsersRepository
}

func NewAuthService(dependencies AuthServiceDependencies) *AuthService {
	return &AuthService{
		UsersRespository: dependencies.UsersRepository,
	}
}

func (service *AuthService) RegisterUser(user *database.User) (*database.User, error) {
	user, err := service.UsersRespository.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
