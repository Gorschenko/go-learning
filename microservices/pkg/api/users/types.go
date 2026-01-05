package users_api

type UserFiltersDto struct {
	ID    int    `json:"ID" query:"ID" url:"ID"`
	Email string `json:"email" query:"email" url:"email"`
}
