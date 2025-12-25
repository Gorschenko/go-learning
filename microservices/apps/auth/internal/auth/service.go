package auth

import (
	"errors"
	"pkg/database"
	"pkg/static"
)

func NewAuthService(dependencies AuthServiceDependencies) *AuthService {
	return &AuthService{
		UsersRespository: dependencies.UsersRepository,
	}
}

func (s *AuthService) RegisterUser(user *database.User) (*database.User, error) {
	existedUser, _ := s.UsersRespository.FindByEmail(user.Email)

	if existedUser != nil {
		return nil, errors.New(static.ErrorUserAlreadyExists)
	}

	user, err := s.UsersRespository.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
