package users

import "pkg/cache"

type UsersHandlerDependencies struct {
	UsersService *UsersService
}

type UsersHandler struct {
	UsersService *UsersService
}

type UsersServiceDependencies struct {
	UsersRepository      *UsersRepository
	CacheUsersRepository *cache.CacheUsersRepository
}

type UsersService struct {
	UsersRepository      *UsersRepository
	CacheUsersRepository *cache.CacheUsersRepository
}

type UserFilters struct {
	ID    int
	Email string
}
