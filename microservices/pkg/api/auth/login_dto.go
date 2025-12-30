package auth_api

import "time"

const (
	LoginPath   = "/auth/login"
	LoginMethod = "POST"
)

type LoginRequestBodyDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseBodyDto struct {
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"expirationTime"`
}
