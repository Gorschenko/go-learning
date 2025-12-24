package auth

import (
	"net/http"
	"pkg/api"
	auth_api "pkg/api/auth"
	"pkg/database"
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

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := r.Context().Value(static.ContextBodyKey).(auth_api.RegisterBodyRequestDto)

		// fmt.Printf("Body: %+v\n", body)
		// params, _ := r.Context().Value(static.ContextParamsKey).(UserPathParams)
		// fmt.Printf("Params: %+v\n", params)

		user := database.User{
			Email:    body.Email,
			Password: body.Password,
			Name:     body.Name,
		}

		createdUser, err := handler.AuthService.RegisterUser(&user)

		if err != nil && err.Error() == static.ErrorUserAlreadyExists {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := auth_api.RegisterBodyResponseDto{
			ID: int(createdUser.ID),
		}

		api.SendJSON(w, response, http.StatusOK)
	}
}
