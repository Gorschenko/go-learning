package auth

import (
	"net/http"
	"pkg/api"
	auth_api "pkg/api/auth"
	"pkg/database"
	"pkg/middlewares"
	"pkg/static"
	"time"
)

func NewAuthHandler(router *http.ServeMux, dependencies AuthHandlerDependencies) {
	handler := &AuthHandler{
		AuthService: dependencies.AuthService,
	}

	authResiterURL := auth_api.AuthRegisterMethod + " " + auth_api.AuthRegisterPath
	router.Handle(
		authResiterURL,
		middlewares.ValidateBody[auth_api.RegisterRequestBodyDto](handler.Register()),
	)

	authLoginURL := auth_api.AuthLoginMethod + " " + auth_api.AuthLoginPath
	router.Handle(
		authLoginURL,
		middlewares.ValidateBody[auth_api.LoginRequestBodyDto](handler.Login()),
	)
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := r.Context().Value(static.ContextBodyKey).(auth_api.RegisterRequestBodyDto)

		// fmt.Printf("Body: %+v\n", body)
		// params, _ := r.Context().Value(static.ContextParamsKey).(UserPathParams)
		// fmt.Printf("Params: %+v\n", params)

		user := database.User{
			Email:    body.Email,
			Password: body.Password,
			Name:     body.Name,
		}

		token, err := h.AuthService.RegisterUser(&user)

		if err != nil && err.Error() == static.ErrorUserAlreadyExists {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := auth_api.RegisterResponseBodyDto{
			Token:          token.Token,
			ExpirationTime: token.ExpirationTime,
		}

		api.SendJSON(w, response, http.StatusOK)
	}
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(6 * time.Second)
		body, _ := r.Context().Value(static.ContextBodyKey).(auth_api.LoginRequestBodyDto)

		token, err := h.AuthService.LoginUser(body.Email, body.Password)

		if err != nil && err.Error() == static.ErrorUserNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if err != nil && err.Error() == static.ErrorInvalidPassowrd {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := auth_api.LoginResponseBodyDto{
			Token:          token.Token,
			ExpirationTime: token.ExpirationTime,
		}

		api.SendJSON(w, response, http.StatusOK)
	}
}
