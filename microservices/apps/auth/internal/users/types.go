package users

type UsersHandlerDependencies struct {
	UsersService *UsersService
}

type UsersHandler struct {
	UsersService *UsersService
}

type UsersServiceDependencies struct {
	UsersRepository *UsersRepository
}

type UsersService struct {
	UsersRepository *UsersRepository
}

type GetOneUserFilters struct {
	ID    int
	Email string
}
