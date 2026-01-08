package auth

import (
	"auth/internal/users"
	"pkg/configs"
)

type AuthHandlerDependencies struct {
	AuthService *AuthService
}

type AuthHandler struct {
	authService *AuthService
}

type AuthServiceDependencies struct {
	Config          *configs.Config
	UsersService    *users.UsersService
	UsersRepository *users.UsersRepository
}

type AuthService struct {
	config           *configs.Config
	usersService     *users.UsersService
	usersRespository *users.UsersRepository
}
