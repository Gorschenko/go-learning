package auth_api

import (
	"pkg/database"
)

const (
	RegisterPath   = "/auth/register"
	RegisterMethod = "POST"
)

type RegisterRequestBodyDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type RegisterResponseBodyDto struct {
	User *database.User `json:"user"`
}
