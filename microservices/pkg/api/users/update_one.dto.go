package users_api

const (
	UpdateOnePath   = "/users"
	UpdateOneMethod = "PATCH"
)

type UpdateOneRequestBodyDto struct {
	Filters UserFiltersDto `json:"filters" validate:"required"`
	Update  UserUpdateDto  `json:"update" validate:"required"`
}

type UpdateOneResponseBodyDto struct {
	Count int `json:"count"`
}
