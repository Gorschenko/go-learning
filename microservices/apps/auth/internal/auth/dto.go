package auth

const (
	AuthRegisterPath = "POST /auth/register"
)

type RegisterBodyRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
