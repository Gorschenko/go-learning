package users_api

const (
	DeleteOnePath   = "/users"
	DeleteOneMethod = "DELETE"
)

type DeleteOneResponseBodyDto struct {
	Count int `json:"count"`
}
