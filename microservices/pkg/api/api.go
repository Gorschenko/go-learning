package api

import (
	"net/http"
	"pkg/configs"
	"time"
)

type HttpApiDependencies struct {
	Config *configs.Config
}

type HttpApi struct {
	Client *http.Client
}

func NewHttpApi(dependencies *HttpApiDependencies) *HttpApi {
	client := &http.Client{
		Timeout: time.Duration(dependencies.Config.Software.Api.TimeoutSec) * time.Second,
	}
	return &HttpApi{
		Client: client,
	}
}
