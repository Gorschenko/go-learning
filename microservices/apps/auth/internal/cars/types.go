package cars

import "auth/internal/users"

type CarsHandlerDependencies struct {
	CarsService *CarsService
}

type CarsHandler struct {
	CarsService *CarsService
}

type CarsServiceDependencies struct {
	CarsRepository  *CarsRepository
	UsersRepository *users.UsersRepository
}

type CarsService struct {
	CarsRepository  *CarsRepository
	UsersRepository *users.UsersRepository
}
