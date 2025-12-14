package auth

import (
	"fmt"
	"net/http"
	"pkg/middlewares"
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

	router.Handle(
		AuthRegisterPath,
		middlewares.ValidateBody[RegisterBodyRequestDto](controller.Register()),
	)

}

func (controller *AuthController) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := r.Context().Value(middlewares.ContextBodyKey).(RegisterBodyRequestDto)

		fmt.Printf("Body: %s", body)
	}
}
