package users_api

import "pkg/database"

const (
	GetOnePath   = "/users"
	GetOneMethod = "GET"
)

type GetOneRequestQueryDto struct {
	UserID int    `json:"userID"`
	Email  string `json:"email"`
}

type GetOneResponseBodyDto struct {
	User *database.User `json:"user"`
}
