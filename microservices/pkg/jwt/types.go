package jwt

import "pkg/configs"

type JWTDependencies struct {
	Config *configs.Config
}

type JWTToken struct {
	Secret   string
	TTLHours int
}

type JWTData struct {
	UserID int
	Email  string
}

type JWT struct {
	Access *JWTToken
}
