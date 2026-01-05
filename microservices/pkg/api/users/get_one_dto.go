package users_api

import (
	"pkg/database"
)

const (
	GetOnePath   = "/users"
	GetOneMethod = "GET"
)

type GetOneResponseBodyDto struct {
	User *database.User `json:"user"`
}
