- Наследование интерфейсов

```
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
    LoginRequest
	Name     string `json:"name" validate:"required"`
}
```