package auth_api

import (
	"pkg/api"
	"pkg/configs"
	"pkg/database"
	"strconv"
)

type AuthApiDependencies struct {
	Config  *configs.Config
	HttpApi *api.HttpApi
}

type AuthApi struct {
	HttpApi *api.HttpApi
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

func (api *AuthApi) RegisterUser(body *database.User) (string, error) {
	url := api.BaseURL + AuthRegisterPath
	response, err := api.HttpApi.Client.
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Execute(AuthRegisterMethod, url)

	if err != nil || response.IsError() {
		return "", err
	}

	return "User", nil
}
