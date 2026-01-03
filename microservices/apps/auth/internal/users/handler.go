package users

import (
	"net/http"
	"pkg/api"
	users_api "pkg/api/users"
	"pkg/middlewares"
	"pkg/static"
)

func NewUsersHandler(router *http.ServeMux, dependencies *UsersHandlerDependencies) {
	handler := &UsersHandler{
		UsersService: dependencies.UsersService,
	}

	getOneURL := users_api.GetOneMethod + " " + users_api.GetOnePath
	router.Handle(
		getOneURL,
		middlewares.ValidateQuery[users_api.GetOneRequestQueryDto](handler.GetOne()),
	)
}

func (h *UsersHandler) GetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := r.Context().Value(static.ContextQueryKey).(users_api.GetOneRequestQueryDto)

		filters := GetOneUserFilters{
			UserID: body.UserID,
			Email:  body.Email,
		}

		user, err := h.UsersService.GetOne(&filters)

		if err != nil {
			api.SendJSONError(w, err)
			return
		}

		response := users_api.GetOneResponseBodyDto{
			User: user,
		}

		api.SendJSON(w, response, http.StatusOK)
	}
}
