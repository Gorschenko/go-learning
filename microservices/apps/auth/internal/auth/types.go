package auth

import (
	"auth/internal/users"
	"pkg/configs"
)

type AuthHandlerDependencies struct {
	AuthService *AuthService
}

type AuthHandler struct {
	AuthService *AuthService
}

type AuthServiceDependencies struct {
	Config          *configs.Config
	UsersService    *users.UsersService
	UsersRepository *users.UsersRepository
}

type AuthService struct {
	Config           *configs.Config
	UsersService     *users.UsersService
	UsersRespository *users.UsersRepository
}
