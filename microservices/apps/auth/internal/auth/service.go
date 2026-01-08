package auth

import (
	"context"
	"errors"
	"pkg/api"
	users_api "pkg/api/users"
	"pkg/database"
	"pkg/jwt"
)

func NewAuthService(dependencies *AuthServiceDependencies) *AuthService {
	return &AuthService{
		Config:           dependencies.Config,
		UsersService:     dependencies.UsersService,
		UsersRespository: dependencies.UsersRepository,
	}
}

func (s *AuthService) RegisterUser(ctx context.Context, user *database.User) (*database.User, error) {
	filters := users_api.UserFiltersDto{
		Email: user.Email,
	}

	existedUser, _ := s.UsersService.GetOne(ctx, &filters)

	if existedUser != nil {
		return nil, errors.New(api.CodeAlreadyExists)
	}

	createdUser, err := s.UsersRespository.Create(user)

	return createdUser, err
}

func (s *AuthService) LoginUser(ctx context.Context, email, password string) (*jwt.JWTToken, error) {
	filters := users_api.UserFiltersDto{
		Email: email,
	}

	existedUser, err := s.UsersService.GetOne(ctx, &filters)

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
