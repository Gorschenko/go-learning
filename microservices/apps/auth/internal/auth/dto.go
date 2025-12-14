package auth

const (
	AuthRegisterPath = "POST /auth/register/{userID}/{action}"
)

type UserPathParams struct {
	UserID int    `param:"userID" validate:"required,gt=0"`
	Action string `param:"action" validate:"required,oneof=view edit delete"`
}

type RegisterBodyRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
