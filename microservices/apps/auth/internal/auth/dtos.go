package auth

type UserPathParams struct {
	UserID int    `param:"userID" validate:"required,gt=0"`
	Action string `param:"action" validate:"required,oneof=view edit delete"`
}
