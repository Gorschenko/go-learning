package users

func NewUsersService(dependencies *UsersServiceDependencies) *UsersService {
	return &UsersService{
		UsersRepository: dependencies.UsersRepository,
	}
}
