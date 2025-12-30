package auth

import (
	"errors"
	"pkg/database"
	"pkg/jwt"
	"pkg/static"
)

func NewAuthService(dependencies *AuthServiceDependencies) *AuthService {
	return &AuthService{
		Config:           dependencies.Config,
		UsersRespository: dependencies.UsersRepository,
	}
}

func (s *AuthService) RegisterUser(user *database.User) (*jwt.JWTToken, error) {
	existedUser, _ := s.UsersRespository.FindByEmail(user.Email)

	if existedUser != nil {
		return nil, errors.New(static.ErrorUserAlreadyExists)
	}

	createdUser, err := s.UsersRespository.Create(user)

	if err != nil {
		return nil, errors.New(static.ErrorInternalServerError)
	}

	payload := jwt.JWTDataToCreate{
		UserID: int(createdUser.ID),
		Email:  createdUser.Email,
	}

	token := jwt.NewJWT(jwt.JWTDependencies{
		Config: s.Config,
	}).Create(payload)

	return token, nil
}

func (s *AuthService) LoginUser(email, password string) (*jwt.JWTToken, error) {
	existedUser, err := s.UsersRespository.FindByEmail(email)

	if err != nil {
		return nil, errors.New(static.ErrorUserNotFound)
	}

	if existedUser.Password != password {
		return nil, errors.New(static.ErrorInvalidPassowrd)
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
