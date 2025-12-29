package jwt

import (
	"pkg/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTDependencies struct {
	Config *configs.Config
}

type JWTAccessConfig struct {
	Secret    string
	ExpiresIn time.Duration
}

type JWT struct {
	AccessConfig *JWTAccessConfig
}

type JWTDataToCreate struct {
	UserID int
	Email  string
}

type JWTClaims struct {
	UserID string
	Email  string
	jwt.RegisteredClaims
}

type JWTToken struct {
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"expirationTime"`
}
