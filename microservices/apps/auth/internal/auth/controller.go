package auth

import (
	"fmt"
	"net/http"
	"pkg/api"
	"pkg/middlewares"
	"pkg/static"
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
		middlewares.ValidateParams[UserPathParams](middlewares.ValidateBody[RegisterBodyRequestDto](controller.Register())),
	)

}

func (controller *AuthController) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := r.Context().Value(static.ContextBodyKey).(RegisterBodyRequestDto)

		fmt.Printf("Body: %+v\n", body)
		// params, _ := r.Context().Value(static.ContextParamsKey).(UserPathParams)
		// fmt.Printf("Params: %+v\n", params)
		api.JSON(w, body, http.StatusOK)
	}
}
