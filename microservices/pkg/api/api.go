package api

import (
	"pkg/configs"
	"time"

	"github.com/go-resty/resty/v2"
)

type HttpApiDependencies struct {
	Config *configs.Config
}

type HttpApi struct {
	Client *resty.Client
}

func NewHttpApi(dependencies *HttpApiDependencies) *HttpApi {
	timeout := time.Duration(dependencies.Config.Software.Api.TimeoutSec) * time.Second
	client := resty.New().SetTimeout(timeout)

	if dependencies.Config.Software.Api.Debug {
		client.SetDebug(true)
	}

	return &HttpApi{
		Client: client,
	}
}
