package auth

import (
	"errors"
	"test/internal/users"
	"test/packages/di"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UsersRepository di.IUsersRepository
}

func NewAuthService(usersRepository di.IUsersRepository) *AuthService {
	return &AuthService{
		UsersRepository: usersRepository,
	}
}

func (service *AuthService) Login(email, password string) (string, error) {
	existedUser, err := service.UsersRepository.FindByEmail(email)

	if err != nil {
		return "", errors.New(ErrorUserNotFound)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))

	if err != nil {
		return "", errors.New(ErrorInvalidPassword)
	}

	return email, nil
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UsersRepository.FindByEmail(email)

	if existedUser != nil {
		return "", errors.New(ErrorUserExist)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &users.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}

	_, err = service.UsersRepository.Create(user)

	if err != nil {
		return "", err
	}

	return user.Email, nil
}
