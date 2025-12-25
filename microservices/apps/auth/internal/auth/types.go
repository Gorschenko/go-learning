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
	UsersRepository *users.UsersRepository
}

type AuthService struct {
	Config           *configs.Config
	UsersRespository *users.UsersRepository
}
