package users_api

import (
	"pkg/database"
)

const (
	GetOnePath   = "/users"
	GetOneMethod = "GET"
)

type GetOneRequestQueryDto struct {
	ID    int    `query:"ID" url:"ID"`
	Email string `query:"email" url:"email"`
}

type GetOneResponseBodyDto struct {
	User *database.User `json:"user"`
}
