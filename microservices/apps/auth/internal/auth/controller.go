package auth

import (
	"net/http"
)

type AuthControllerDependencies struct {
	AuthService *AuthService
}

type AuthController struct {
	AuthService *AuthService
}

func NewAuthController(router *http.ServeMux, dependencies AuthControllerDependencies) {
	controller := &AuthController{
		AuthService: dependencies.AuthService,
	}

	router.HandleFunc("POST /auth/register", controller.Register())
}

func (controller *AuthController) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
