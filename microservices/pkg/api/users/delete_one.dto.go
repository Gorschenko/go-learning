package users_api

const (
	DeleteOnePath   = "/users"
	DeleteOneMethod = "DELETE"
)

type DeleteOneRequestBodyDto struct {
	ID    int    `json:"ID"`
	Email string `json:"email"`
}

type DeleteOneResponseBodyDto struct {
	Count int `json:"count"`
}
