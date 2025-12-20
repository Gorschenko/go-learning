package api

import (
	"pkg/configs"
	"strconv"
)

type AuthApiDependencies struct {
	Config  *configs.Config
	HttpApi *HttpApi
}

type AuthApi struct {
	HttpApi *HttpApi
	BaseURL string
}

func NewAuthApi(dependencies *AuthApiDependencies) *AuthApi {
	protocol := dependencies.Config.Services.Auth.Protocol
	hostname := dependencies.Config.Services.Auth.Host
	port := dependencies.Config.Services.Auth.Port

	baseURL := protocol + "://" + hostname + ":" + strconv.Itoa(port)

	return &AuthApi{
		HttpApi: dependencies.HttpApi,
		BaseURL: baseURL,
	}
}
