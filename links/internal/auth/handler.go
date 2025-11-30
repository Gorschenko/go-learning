package auth

import (
	"net/http"
	"test/configs"
	"test/packages/request"
	"test/packages/response"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, r)

		if err != nil {
			return
		}

		token, err := handler.AuthService.Login(body.Email, body.Password)

		if err != nil {
			response.Json(w, err, http.StatusUnauthorized)
			return
		}

		data := LoginResponse{
			Token: token,
		}
		response.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, r)

		if err != nil {
			return
		}

		email, _ := handler.AuthService.Register(body.Email, body.Password, body.Name)

		response.Json(w, email, http.StatusOK)
	}
}
