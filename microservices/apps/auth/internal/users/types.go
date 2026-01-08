package users

import "pkg/cache"

type UsersHandlerDependencies struct {
	UsersService *UsersService
}

type UsersHandler struct {
	usersService *UsersService
}

type UsersServiceDependencies struct {
	UsersRepository      *UsersRepository
	CacheUsersRepository *cache.CacheUsersRepository
}

type UsersService struct {
	usersRepository      *UsersRepository
	cacheUsersRepository *cache.CacheUsersRepository
}
