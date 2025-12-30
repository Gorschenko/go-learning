package auth

type TestUserPathParams struct {
	UserID int    `param:"userID" validate:"required,gt=0"`
	Action string `param:"action" validate:"required,oneof=view edit delete"`
}

type TestUserQuery struct {
	UserID int    `query:"userID" validate:"required,gt=0"`
	Action string `query:"action" validate:"required,oneof=view edit delete"`
}
