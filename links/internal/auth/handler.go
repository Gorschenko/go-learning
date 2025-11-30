package auth

import (
	"net/http"
	"test/configs"
	"test/packages/jwt"
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

		email, err := handler.AuthService.Login(body.Email, body.Password)

		if err != nil {
			response.Json(w, err, http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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

		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)

		if err != nil {
			response.Json(w, err, http.StatusBadRequest)
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data := RegisterResponse{
			Token: token,
		}

		response.Json(w, data, http.StatusOK)
	}
}
