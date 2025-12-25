package auth

import "auth/internal/users"

type AuthHandlerDependencies struct {
	AuthService *AuthService
}

type AuthHandler struct {
	AuthService *AuthService
}

type AuthServiceDependencies struct {
	UsersRepository *users.UsersRepository
}

type AuthService struct {
	UsersRespository *users.UsersRepository
}
