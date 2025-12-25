package auth

import (
	"errors"
	"pkg/database"
	"pkg/jwt"
	"pkg/static"
	"time"
)

func NewAuthService(dependencies AuthServiceDependencies) *AuthService {
	return &AuthService{
		Config:           dependencies.Config,
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

func (s *AuthService) LoginUser(email, password string) (string, time.Time, error) {
	existedUser, err := s.UsersRespository.FindByEmail(email)

	if err != nil {
		return "", time.Time{}, errors.New(static.ErrorUserNotFound)
	}

	if existedUser.Password != password {
		return "", time.Time{}, errors.New(static.ErrorInvalidPassowrd)
	}

	payload := jwt.JWTPayload{
		UserID: int(existedUser.ID),
		Email:  email,
	}

	token, expirationTime := jwt.NewJWT(jwt.JWTDependencies{
		Config: s.Config,
	}).Create(payload)

	return token, expirationTime, err
}
