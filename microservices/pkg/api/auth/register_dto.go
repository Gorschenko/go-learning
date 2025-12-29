package auth_api

import "time"

const (
	AuthRegisterPath   = "/auth/register"
	AuthRegisterMethod = "POST"
)

type RegisterRequestBodyDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type RegisterResponseBodyDto struct {
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"expirationTime"`
}
