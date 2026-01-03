package auth

import (
	"auth/internal/users"
	"errors"
	"pkg/api"
	"pkg/database"
	"pkg/jwt"
)

func NewAuthService(dependencies *AuthServiceDependencies) *AuthService {
	return &AuthService{
		Config:           dependencies.Config,
		UsersRespository: dependencies.UsersRepository,
	}
}

func (s *AuthService) RegisterUser(user *database.User) (*database.User, error) {
	filters := users.UserFilters{
		Email: user.Email,
	}
	existedUser, _ := s.UsersRespository.GetOne(&filters)

	if existedUser != nil {
		return nil, errors.New(api.CodeAlreadyExists)
	}

	createdUser, err := s.UsersRespository.Create(user)

	return createdUser, err
}

func (s *AuthService) LoginUser(email, password string) (*jwt.JWTToken, error) {
	filters := users.UserFilters{
		Email: email,
	}

	existedUser, err := s.UsersRespository.GetOne(&filters)

	if err != nil {
		return nil, err
	}

	if existedUser == nil {
		return nil, errors.New(api.CodeNotFound)
	}

	if existedUser.Password != password {
		return nil, errors.New(api.CodeUnauthorized)
	}

	payload := jwt.JWTDataToCreate{
		UserID: int(existedUser.ID),
		Email:  email,
	}

	token := jwt.NewJWT(jwt.JWTDependencies{
		Config: s.Config,
	}).Create(payload)

	return token, err
}
