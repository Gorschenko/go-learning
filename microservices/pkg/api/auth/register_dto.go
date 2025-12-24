package auth_api

const (
	AuthRegisterPath   = "/auth/register"
	AuthRegisterMethod = "POST"
)

type RegisterBodyRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type RegisterBodyResponseDto struct {
	ID int `json:"ID"`
}
