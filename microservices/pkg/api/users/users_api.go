package users_api

import (
	"pkg/api"
	"pkg/configs"
	"strconv"
)

type UsersApiDependencies struct {
	HttpApi *api.HttpApi
}

type UsersApi struct {
	httpApi *api.HttpApi
	baseURL string
}

func NewUsersApi(dependencies *UsersApiDependencies) *UsersApi {
	return &UsersApi{
		httpApi: dependencies.HttpApi,
		baseURL: "",
	}
}

func (a *UsersApi) SetBaseURLByConfig(config *configs.Config) *UsersApi {
	protocol := config.Services.Auth.Protocol
	hostname := config.Services.Auth.Host
	port := config.Services.Auth.Port

	baseURL := protocol + "://" + hostname + ":" + strconv.Itoa(port)

	a.baseURL = baseURL

	return a
}

func (a *UsersApi) SetBaseURL(URL string) *UsersApi {
	a.baseURL = URL
	return a
}

func (a *UsersApi) DeleteOneUser(body *UserFiltersDto) (*DeleteOneResponseBodyDto, error) {
	URL := a.baseURL + DeleteOnePath

	var result DeleteOneResponseBodyDto

	response, err := a.httpApi.Client.
		R().
		SetBody(body).
		SetResult(&result).
		Execute(DeleteOneMethod, URL)

	if err != nil || response.IsError() {
		return nil, err
	}

	return &result, nil
}
