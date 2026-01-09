package users

import (
	"pkg/cache"
	"pkg/database"
)

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

type UsersRepository struct {
	database *database.Db
}
