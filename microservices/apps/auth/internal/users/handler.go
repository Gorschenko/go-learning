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
		middlewares.ValidateQuery[users_api.UserFiltersDto](handler.GetOne()),
	)

	deleteOneURL := users_api.DeleteOneMethod + " " + users_api.DeleteOnePath
	router.Handle(
		deleteOneURL,
		middlewares.ValidateBody[users_api.UserFiltersDto](handler.DeleteOne()),
	)

	updateOneURL := users_api.UpdateOneMethod + " " + users_api.UpdateOnePath
	router.Handle(
		updateOneURL,
		middlewares.ValidateBody[users_api.UpdateOneRequestBodyDto](handler.UpdateOne()),
	)
}

func (h *UsersHandler) GetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, _ := ctx.Value(static.ContextQueryKey).(users_api.UserFiltersDto)

		user, err := h.UsersService.GetOne(ctx, &body)

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

func (h *UsersHandler) UpdateOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, _ := ctx.Value(static.ContextBodyKey).(users_api.UpdateOneRequestBodyDto)

		count, err := h.UsersService.UpdateOne(ctx, &body.Filters, &body.Update)

		if err != nil {
			api.SendJSONError(w, err)
			return
		}

		response := users_api.UpdateOneResponseBodyDto{
			Count: count,
		}

		api.SendJSON(w, response, http.StatusOK)
	}
}

func (h *UsersHandler) DeleteOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, _ := ctx.Value(static.ContextBodyKey).(users_api.UserFiltersDto)

		count, err := h.UsersService.DeleteOne(ctx, &body)

		if err != nil {
			api.SendJSONError(w, err)
			return
		}

		response := users_api.DeleteOneResponseBodyDto{
			Count: count,
		}

		api.SendJSON(w, response, http.StatusOK)
	}
}
