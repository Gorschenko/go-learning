package auth

import (
	"fmt"
	"net/http"
	"pkg/api"
	auth_api "pkg/api/auth"
	"pkg/middlewares"
	"pkg/static"
)

type AuthHandlerDependencies struct {
	AuthService *AuthService
}

type AuthHandler struct {
	AuthService *AuthService
}

func NewAuthHandler(router *http.ServeMux, dependencies AuthHandlerDependencies) {
	handler := &AuthHandler{
		AuthService: dependencies.AuthService,
	}

	authResiterURL := auth_api.AuthRegisterMethod + " " + auth_api.AuthRegisterPath
	router.Handle(
		authResiterURL,
		middlewares.ValidateBody[auth_api.RegisterBodyRequestDto](handler.Register()),
	)

}

func (controller *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := r.Context().Value(static.ContextBodyKey).(auth_api.RegisterBodyRequestDto)

		fmt.Printf("Body: %+v\n", body)
		// params, _ := r.Context().Value(static.ContextParamsKey).(UserPathParams)
		// fmt.Printf("Params: %+v\n", params)
		api.JSON(w, body, http.StatusOK)
	}
}
