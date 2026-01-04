package auth

import (
	"net/http"
	"pkg/api"
	auth_api "pkg/api/auth"
	"pkg/database"
	"pkg/middlewares"
	"pkg/static"
)

func NewAuthHandler(router *http.ServeMux, dependencies *AuthHandlerDependencies) {
	handler := &AuthHandler{
		AuthService: dependencies.AuthService,
	}

	registerURL := auth_api.RegisterMethod + " " + auth_api.RegisterPath
	router.Handle(
		registerURL,
		middlewares.ValidateBody[auth_api.RegisterRequestBodyDto](handler.RegisterUser()),
	)

	loginURL := auth_api.LoginMethod + " " + auth_api.LoginPath
	router.Handle(
		loginURL,
		middlewares.ValidateBody[auth_api.LoginRequestBodyDto](handler.LoginUser()),
	)
}

func (h *AuthHandler) RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := r.Context().Value(static.ContextBodyKey).(auth_api.RegisterRequestBodyDto)

		user := database.User{
			Email:    body.Email,
			Password: body.Password,
			Name:     body.Name,
		}

		createdUser, err := h.AuthService.RegisterUser(&user)

		if err != nil {
			api.SendJSONError(w, err)
			return
		}

		response := auth_api.RegisterResponseBodyDto{
			User: createdUser,
		}

		api.SendJSON(w, response, http.StatusOK)
	}
}

func (h *AuthHandler) LoginUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := r.Context().Value(static.ContextBodyKey).(auth_api.LoginRequestBodyDto)

		token, err := h.AuthService.LoginUser(body.Email, body.Password)

		if err != nil {
			api.SendJSONError(w, err)
			return
		}

		response := auth_api.LoginResponseBodyDto{
			Token:          token.Token,
			ExpirationTime: token.ExpirationTime,
		}

		api.SendJSON(w, response, http.StatusOK)
	}
}
