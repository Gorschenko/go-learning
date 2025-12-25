package auth_api

import "time"

const (
	AuthLoginPath   = "/auth/login"
	AuthLoginMethod = "POST"
)

type LoginBodyRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginBodyResponseDto struct {
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"expirationTime"`
}
