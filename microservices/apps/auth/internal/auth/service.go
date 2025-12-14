package auth

import "auth/internal/users"

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
