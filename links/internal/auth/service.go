package auth

import (
	"errors"
	"test/internal/users"
)

type AuthService struct {
	UsersRepository *users.UsersRepository
}

func NewAuthService(usersRepository *users.UsersRepository) *AuthService {
	return &AuthService{
		UsersRepository: usersRepository,
	}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UsersRepository.FindByEmail(email)

	if existedUser != nil {
		return "", errors.New(ErrorUserExist)
	}

	user := &users.User{
		Email:    email,
		Password: "",
		Name:     name,
	}

	_, err := service.UsersRepository.Create(user)

	if err != nil {
		return "", err
	}

	return user.Email, nil
}
