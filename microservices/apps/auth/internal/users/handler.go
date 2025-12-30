package users

import (
	"net/http"
	users_api "pkg/api/users"
	"pkg/middlewares"
)

func NewUsersHandler(router *http.ServeMux, dependencies *UsersHandlerDependencies) {
	handler := &UsersHandler{
		UsersService: dependencies.UsersService,
	}

	getOneURL := users_api.GetOneMethod + " " + users_api.GetOnePath
	router.Handle(
		getOneURL,
		middlewares.ValidateBody[users_api.GetOneRequestQueryDto](handler.GetOne()),
	)
}

func (h *UsersHandler) GetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
