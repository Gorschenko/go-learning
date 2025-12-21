package api

import (
	"encoding/json"
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

func (api *AuthApi) RegisterUser(body *database.User) (*database.User, error) {
	url := api.BaseURL + "/auth"
	response, err := api.HttpApi.Client.
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}

	var user database.User

	json.Unmarshal(response.Body(), &user)

	return &user, err
}
