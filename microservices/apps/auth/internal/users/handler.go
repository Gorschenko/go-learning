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

	deleteOneURL := users_api.DeleteOneMethod + " " + users_api.DeleteOnePath
	router.Handle(
		deleteOneURL,
		middlewares.ValidateBody[users_api.DeleteOneRequestBodyDto](handler.DeleteOne()),
	)
}

func (h *UsersHandler) GetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := r.Context().Value(static.ContextQueryKey).(users_api.GetOneRequestQueryDto)

		filters := UserFilters{
			ID:    body.ID,
			Email: body.Email,
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

func (h *UsersHandler) DeleteOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := r.Context().Value(static.ContextBodyKey).(users_api.DeleteOneRequestBodyDto)

		filters := UserFilters{
			ID:    body.ID,
			Email: body.Email,
		}

		count, err := h.UsersService.DeleteOne(&filters)

		if err != nil {
			api.SendJSONError(w, err)
		}

		response := users_api.DeleteOneResponseBodyDto{
			Count: count,
		}

		api.SendJSON(w, response, http.StatusOK)
	}
}
