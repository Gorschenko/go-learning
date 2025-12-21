package api

import (
	"errors"
	"log"
	"pkg/configs"
	"pkg/database"
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

func (api *AuthApi) RegisterUser(body *database.User) (string, error) {
	url := api.BaseURL + "/auth/register"
	response, err := api.HttpApi.Client.
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)
	log.Printf("RESPONSE: %s", response)

	if err != nil || response.IsError() {
		return "", errors.New("Bad external request")
	}

	return "User", nil
}
